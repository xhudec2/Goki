package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const CARD_DELIMITER = "\x1f"

type ID int

type Gettable interface {
	Deck | Card | Note
	GetID() ID
}

func (d Deck) GetID() ID {
	return ID(d.ID)
}

func (c Card) GetID() ID {
	return ID(c.ID)
}

func (n Note) GetID() ID {
	return ID(n.ID)
}

type Table[T Gettable] map[ID]T

// This function is nearly the same as the one in gorm.sqlite,
// however, it uses a collating function defined in collation/collation.go
func Open(dsn string) gorm.Dialector {
	return &sqlite.Dialector{DSN: dsn, DriverName: "sqlite_unicase"}
}

func OpenDB(filepath string) (db *gorm.DB, err error) {
	db, err = gorm.Open(Open(filepath))
	if err != nil {
		log.Fatal("Error opening DB:", err)
		return
	}
	return
}

func DBGetter[T Gettable](db *gorm.DB, data *Table[T]) {
	var listData []T
	db.Find(&listData)

	for i := 0; i < len(listData); i++ {
		(*data)[listData[i].GetID()] = listData[i]
	}
}
