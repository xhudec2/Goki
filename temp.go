package main

//
// // TEMP file, only for testing
//
// import (
// 	"bufio"
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"project/database"
// 	"project/scheduler"
// )
//
// func Printall(decks database.Decks, db *sql.DB) {
// 	var all_cards []database.Cards = make([]database.Cards, 0, 4)
// 	for id, name := range decks {
// 		fmt.Printf("Id: %s, Name: %s\n", id, name)
// 		card_ids, cards, err := database.Get_Cards(id, db)
// 		if err != nil {
// 			return
// 		}
// 		database.Get_notes(card_ids, cards, db)
// 		all_cards = append(all_cards, cards)
// 	}
// 	for _, deck := range all_cards {
// 		for id, card := range deck {
// 			fmt.Println(id, card)
// 		}
// 	}
// }
//
// func print_deck(deck_id database.Id, db *sql.DB) {
// 	card_ids, cards, err := database.Get_Cards(deck_id, db)
// 	if err != nil {
// 		return
// 	}
// 	err = database.Get_notes(card_ids, cards, db)
// 	if err != nil {
// 		return
// 	}
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for _, card := range cards {
// 		if card.Back == "" || card.Front == "" {
// 			continue
// 		}
// 		fmt.Printf("%s\n", card.Front)
// 		scanner.Scan()
// 		fmt.Printf("> %s\n", card.Back)
// 		scanner.Scan()
// 		text := scanner.Text()
// 		if text == "exit" {
// 			return
// 		}
// 	}
// }
//
// func Test_db(db *sql.DB) {
// 	decks, err := database.Get_decks(db)
// 	if err != nil {
// 		return
// 	}
// 	delete(decks, "1")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for id, name := range decks {
// 		fmt.Printf("Do you want to study %s ?\n> ", name)
// 		scanner.Scan()
// 		response := scanner.Text()
// 		switch response {
// 		case "y", "yes", "Y", "YES":
// 			print_deck(id, db)
// 		case "e", "E", "exit", "EXIT":
// 			return
// 		}
// 	}
// }
//
// func get_deck_id(db *sql.DB) (did database.Id, err error) {
// 	decks, err := database.Get_decks(db)
// 	if err != nil {
// 		return
// 	}
// 	delete(decks, "1")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for id, name := range decks {
// 		fmt.Printf("Do you want to study %s ?\n> ", name)
// 		scanner.Scan()
// 		response := scanner.Text()
// 		switch response {
// 		case "y", "yes", "Y", "YES":
// 			return id, nil
// 		}
// 	}
// 	return
// }
//
// func Test_Scheduler() {
// 	db, err := database.Open_db("Anki2/User 1/collection.anki2")
// 	if err != nil {
// 		return
// 	}
// 	defer db.Close()
// 	did, err := get_deck_id(db)
// 	if err != nil {
// 		return
// 	}
// 	card_ids, cards, err := database.Get_Cards(did, db)
// 	if err != nil {
// 		return
// 	}
// 	err = database.Get_notes(card_ids, cards, db)
// 	if err != nil {
// 		return
// 	}
// 	qs := scheduler.Scheduler_init()
// 	err = qs.Fill_scheduler(card_ids, db)
// 	if err != nil {
// 		return
// 	}
// 	err = qs.Study(cards, db)
// 	if err != nil {
// 		return
// 	}
// }

// // Returns at most how many reviews are possible in the current day
// // No more than leftToday reviews of a card can be done in a single day
// func leftToday(delays []int) int {
// 	now := int(time.Now().Unix())
// 	midnight := now + DAY - (now % DAY)
// 	left := 0
// 	for _, delay := range delays {
// 		now += delay * MINUTE
// 		if midnight <= now {
// 			left += 1
// 		}
// 	}
// 	return left
// }
