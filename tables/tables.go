package tables

import (
	"log"
	"time"

	"gorm.io/gorm"
)

const SECONDS_IN_DAY = 86400

type ID uint64

// https://github.com/ankidroid/Anki-Android/wiki/Database-Structure

// Most of these are unused, only here not to break the code

type Col struct {
	ID     ID `gorm:"primaryKey;autoCreateTime:milli"`
	Crt    int
	Mod    int `gorm:"autoUpdateTime"`
	Scm    int
	Ver    int
	Dty    int
	Usn    int
	Ls     int
	Conf   string
	Models string
	Decks  string
	Dconf  string
	Tags   string
}

type Deck struct {
	ID        ID `gorm:"primaryKey;autoCreateTime:milli"`
	Name      string
	MtimeSecs int
	Usn       int
	Common    []byte
	Kind      []byte
}

func (d Deck) GetID() ID {
	return ID(d.ID)
}

func TodayRelative(db *gorm.DB) int {
	col := Col{}
	if err := db.Table("col").Take(&col).Error; err != nil {
		log.Fatal(err)
	}
	return (int(time.Now().Unix()) - col.Crt) / SECONDS_IN_DAY
}
