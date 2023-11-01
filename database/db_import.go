package database

import (
	"log"

	"gorm.io/gorm"
)

const VER = 18

func ImportDB(source *gorm.DB, target *gorm.DB) error {
	var decks []Deck
	if err := source.Table("decks").Find(&decks).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := target.Table("decks").Create(&decks).Error; err != nil {
		log.Fatal(err)
		return err
	}

	var cards []Card
	if err := source.Table("cards").Find(&cards).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := target.Table("cards").Create(&cards).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
