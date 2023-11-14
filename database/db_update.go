package database

import (
	"fmt"
	. "project/tables"

	"gorm.io/gorm"
)

type Attr_tuple struct {
	Attr_name string
	New_val   string
}

type UpdatedAttributes []Attr_tuple

func Insert_card(db *gorm.DB, card Card) (err error) {
	result := db.Create(&card)
	if result.Error != nil {
		fmt.Println("Error occurred while inserting card: ", result.Error)
		return result.Error
	}
	return nil
}

func UpdateCard(cardID int, db *gorm.DB, attrs []UpdatedAttributes) error {
	result := db.Model(&Card{}).Where("id = ?", cardID).Updates(attrs)
	if result.Error != nil {
		fmt.Printf("Error updating card %d, err: %v", cardID, result.Error)
		return result.Error
	}
	return nil
}

func DeleteCard(cardID int, db *gorm.DB) error {
	result := db.Delete(&Card{}, cardID)
	if result.Error != nil {
		fmt.Printf("Error deleting note %d, err: %v", cardID, result.Error)
		return result.Error
	}
	return nil
}

func DeleteNote(noteID int, db *gorm.DB) error {
	result := db.Delete(&Note{}, noteID)
	if result.Error != nil {
		fmt.Printf("Error deleting note %d, err: %v", noteID, result.Error)
		return result.Error
	}
	return nil
}
