package database

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

func StudyCard(card *Card, db *gorm.DB, flds *map[ID]StudyNote) (err error) {
	// TODO: unfinished, only prints the card for now
	// It should also change the db based on the "grade" given to the card while studying
	// return the grade as well to know where to store the card next

	go card.UpdateCard(db)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(">", (*flds)[card.ID].Front)
	scanner.Scan()
	_ = scanner.Text()
	fmt.Println((*flds)[card.ID].Back)
	fmt.Print("Grade: ")
	scanner.Scan()
	grade := scanner.Text()
	fmt.Println(grade)

	// updateCard

	return nil
}

func (card *Card) UpdateCard(db *gorm.DB) {
	card.Reps = 1
	db.Model(Card{}).Where("id = ?", card.ID).Updates(Card{Reps: card.Reps, Factor: 2500})
}

func GetNids(cardIDs []ID, db *gorm.DB, nids *[]ID) (err error) {

	err = db.Model(Card{}).Select("nid").Find(nids, cardIDs).Error
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)

	}
	return
}

func GetFlds(cardIDs []ID, db *gorm.DB, flds *map[ID]StudyNote) (err error) {

	type loader struct {
		ID   ID
		Flds string
	}
	load := make([]loader, len(cardIDs))

	err = db.Table("notes").
		Select("notes.id, notes.flds").
		Joins("join cards on cards.nid = notes.id").
		Find(&load, cardIDs).Error

	for _, loaded := range load {
		splitted := strings.Split(loaded.Flds, CARD_DELIMITER)
		(*flds)[loaded.ID] = StudyNote{splitted[0], splitted[1]}
	}
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	return
}
