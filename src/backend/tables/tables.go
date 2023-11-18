package tables

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type ID uint64

// https://github.com/ankidroid/Anki-Android/wiki/Database-Structure

// Most of these are unused, only here not to break the code

// it should have more fields, these are good for now

type Col struct {
	ID     ID  `gorm:"primaryKey;autoCreateTime:milli"`
	Crt    int `gorm:"primaryKey;autoCreateTime:milli"`
	Mod    int `gorm:"autoUpdateTime"`
	Scm    int
	Ver    int `gorm:"default:18"`
	Dty    int
	Usn    int
	Ls     int
	Conf   string
	Models string
	Decks  string
	Dconf  string
	Tags   string
}

func (c Col) GetID() ID {
	return ID(c.ID)
}

func TodayRelative(db *gorm.DB) int {
	col := Col{}
	if err := db.Table("col").Take(&col).Error; err != nil {
		log.Fatal(err)
	}
	return (int(time.Now().Unix()) - col.Crt) / DAY
}
