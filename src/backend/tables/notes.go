package tables

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Note struct {
	ID    ID `gorm:"primaryKey;autoCreateTime:milli"`
	Guid  string
	Mid   int
	Mod   int `gorm:"autoUpdateTime"`
	Usn   int
	Tags  string
	Flds  string
	Sfld  string
	Csum  int
	Flags int
	Data  string
}

type NoteFlds struct {
	Nid  ID
	Flds string
}

type StudyNote struct {
	Front string
	Back  string
}

func (n Note) GetID() ID {
	return ID(n.ID)
}

func GetNoteIDs(cardIDs []ID, db *gorm.DB, nids *[]ID) (err error) {

	err = db.Model(Card{}).Select("nid").Find(nids, cardIDs).Error
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	return
}

const CARD_DELIMITER = "\x1f"

func GetFlds(cardIDs *[]ID, db *gorm.DB, flds *map[ID]StudyNote) (err error) {

	type loader struct {
		ID   ID
		Flds string
	}
	load := make([]loader, len(*cardIDs))

	err = db.Table("notes").
		Select("notes.id, notes.flds").
		Joins("join cards on cards.nid = notes.id", *cardIDs).
		Find(&load).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, loaded := range load {
		splitted := strings.Split(loaded.Flds, CARD_DELIMITER)
		(*flds)[loaded.ID] = StudyNote{splitted[0], splitted[1]}
	}
	return
}
