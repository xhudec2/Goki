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
type newConf struct {
	Delays []int
	Ints   []int
	Factor int
}

type reviewStruct struct {
	HardFactor float64
	EasyFactor float64
}

type lapseStruct struct {
	Delays []int
}

type Config struct {
	New    newConf
	Review reviewStruct
	Lapse  lapseStruct
}

var CONFIG = Config{
	New: newConf{
		Delays: []int{1, 10},
		Ints:   []int{1, 4},
		Factor: 2500,
	},
	Review: reviewStruct{
		HardFactor: 1.2,
		EasyFactor: 2.5,
	},
	Lapse: lapseStruct{
		Delays: []int{10},
	},
}

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
	return (int(time.Now().Unix()) - col.Crt) / DAY
}
