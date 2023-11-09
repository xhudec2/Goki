package database

import (
	"log"

	. "project/tables"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gettable interface {
	GetID() ID
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

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DBGetter[T Gettable](db *gorm.DB, data *Table[T], where string) error {
	var listData []T
	if where == "" {
		if err := db.Find(&listData).Error; err != nil {
			log.Fatal(err)
			return err
		}
	} else {
		if err := db.Where(where).Find(&listData).Error; err != nil {
			log.Fatal(err)
			return err
		}
	}

	for i := 0; i < len(listData); i++ {
		(*data)[listData[i].GetID()] = listData[i]
	}
	return nil
}
