package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"project/database"
)

// func printall(decks database.Decks, db *sql.DB) {
// 	var all_cards []database.Cards = make([]database.Cards, 0, 4)
// 	for id, name := range decks {
// 		fmt.Printf("Id: %s, Name: %s\n", id, name)
// 		cards, err := database.Get_Cards(id, db)
// 		if err != nil {
// 			return
// 		}
// 		database.Get_notes(cards, db)
// 		all_cards = append(all_cards, cards)
// 	}
// 	for _, deck := range all_cards {
// 		for id, card := range deck {
// 			fmt.Println(id, card)
// 		}
// 	}
// }

func print_deck(deck_id database.Id, db *sql.DB) {
	cards, err := database.Get_Cards(deck_id, db)
	if err != nil {
		return
	}
	err = database.Get_notes(cards, db)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for _, card := range cards {
		if card.Back == "" || card.Front == "" {
			continue
		}
		fmt.Printf("%s\n", card.Front)
		scanner.Scan()
		fmt.Printf("> %s\n", card.Back)
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			return
		}
	}
}

func main() {

	db, err := database.Open_db("Anki2/User 1/collection.anki2")
	if err != nil {
		return
	}
	defer db.Close()
	decks, err := database.Get_decks(db)
	if err != nil {
		return
	}
	delete(decks, "1")
	scanner := bufio.NewScanner(os.Stdin)
	for id, name := range decks {
		fmt.Printf("Do you want to study %s ?\n> ", name)
		scanner.Scan()
		response := scanner.Text()
		switch response {
		case "y", "yes", "Y", "YES":
			print_deck(id, db)
		case "e", "E", "exit", "EXIT":
			return
		}
	}
}
