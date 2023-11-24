package database

import (
	"fmt"
	. "src/backend/tables"

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

func UpdateCard(cardID ID, db *gorm.DB, attrs []UpdatedAttributes) error {
	result := db.Model(&Card{}).Where("id = ?", cardID).Updates(attrs)
	if result.Error != nil {
		fmt.Printf("Error updating card %d, err: %v", cardID, result.Error)
		return result.Error
	}
	return nil
}

func DeleteCard(cardID ID, db *gorm.DB) error {
	result := db.Delete(&Card{}, cardID)
	if result.Error != nil {
		fmt.Printf("Error deleting note %d, err: %v", cardID, result.Error)
		return result.Error
	}
	return nil
}

func DeleteNote(noteID ID, db *gorm.DB) error {
	result := db.Delete(&Note{}, noteID)
	if result.Error != nil {
		fmt.Printf("Error deleting note %d, err: %v", noteID, result.Error)
		return result.Error
	}
	return nil
}

func AddNewCard(did ID, front string, back string, db *gorm.DB) {
	flds := front + "\x1f" + back
	note := Note{Flds: flds}
	db.Create(&note)

	card := Card{Nid: note.ID, Did: did}
	db.Create(&card)
}

func AddNewDeck(name string, db *gorm.DB) {
	defaultBytes := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	deck := DeckTable{Name: name, Common: defaultBytes, Kind: defaultBytes}
	db.Table("decks").Create(&deck)
}

func DeleteDeck(name string, db *gorm.DB) {
	decks := Table[DeckTable]{}
	DBGetter(db, &decks, "decks", fmt.Sprintf("name = '%s'", name))
	id := decks[0].ID
	db.Table("decks").Where("name = ?", name).Delete(&DeckTable{})

	cards := Table[Card]{}
	DBGetter(db, &cards, "cards", fmt.Sprintf("did = %d", id))
	for _, card := range cards {
		db.Table("notes").Where("nid = ?", card.Nid).Delete(&Note{})
	}
	db.Table("cards").Where("did = ?", id).Delete(&Card{})
}
