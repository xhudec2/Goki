package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const CARD_DELIMITER = "\x1f"

type Name string
type Id string
type Decks map[Id]Name

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
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Error scanning decks rows:", err)
			return
		}
		decks[Id(id)] = Name(name)
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
	for rows.Next() {
		var card_id string
		err = rows.Scan(&card_id)
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
		var card_id, flds string
		err = rows.Scan(&card_id, &flds)
		if err != nil {
			fmt.Println("Error scanning notes rows:", err)
			return
		}
		card := strings.Split(flds, CARD_DELIMITER)
		cards[Id(card_id)] = Card{Front: Front(card[0]), Back: Back(card[1])}
	}
	return
}
