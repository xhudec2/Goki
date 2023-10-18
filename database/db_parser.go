package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const CARD_DELIMITER = "\x1f"
const attributes = "id, nid, did, queue, due, ivl, reps, lapses, left"

type Name string
type Decks map[int]Name

type Back string
type Front string

type Card struct {
	Back  Back
	Front Front
}

type Cards map[Id]Card

func Open_db(filepath string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		fmt.Println("Error opening DB:", err)
		return
	}
	return
}

func Get_decks(db *sql.DB) (decks Decks, err error) {
	rows, err := db.Query("SELECT id, name FROM decks")
	if err != nil {
		fmt.Println("Error querying DB decks:", err)
		return
	}
	defer rows.Close()

	decks = make(Decks, 4)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Error scanning decks rows:", err)
			return
		}
		decks[id] = Name(name)
	}
	return
}

func Get_Cards(deck_id Id, db *sql.DB) (cards Cards, err error) {
	rows, err := db.Query(fmt.Sprintf("SELECT id FROM cards WHERE did == %s", deck_id))
	if err != nil {
		fmt.Println("Error querying DB cards:", err)
		return
	}
	defer rows.Close()

	cards = make(Cards, 32)
	card_ids = make(Card_ids, 0, 32)
	for rows.Next() {
		new_card := Card{}
		err = rows.Scan(
			&new_card.Id, &new_card.Nid, &new_card.Did, &new_card.Q_type, &new_card.Due,
			&new_card.Ivl, &new_card.Reps, &new_card.Lapses, &new_card.Left)

		if err != nil {
			fmt.Println("Error scanning cards rows:", err)
			return
		}
		cards[Id(card_id)] = Card{}
	}

	return
}

func Get_notes(cards Cards, db *sql.DB) (err error) {
	card_ids := make([]string, 0, len(cards))
	for card := range cards {
		card_ids = append(card_ids, string(card))
	}
	rows, err := db.Query(fmt.Sprintf("SELECT id, flds FROM notes WHERE id IN (%s)", strings.Join(card_ids, ",")))
	if err != nil {
		fmt.Println("Error querying DB notes:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var card_id int
		var flds string
		err = rows.Scan(&card_id, &flds)
		if err != nil {
			fmt.Println("Error scanning notes rows:", err)
			return
		}
		//have to rewrite this, again..

		//card := strings.Split(flds, CARD_DELIMITER)
		//fmt.Println(cards[card_id].Field)
		//updated_card := cards[card_id]
		//updated_card.Field = Field{Front: card[0], Back: card[1]}
		//cards[card_id] = updated_card
	}
	return
}
