package main

import (
	"fmt"
	"log"
	. "project/collation"
	. "project/database"
	. "project/scheduler"
	"time"
)

const SECONDS_IN_DAY = 86400

func main() {
	RegisterCollation()
	db, err := OpenDB(MAIN_DB)
	if err != nil {
		return
	}
	defer CloseDB(db)

	cards := make(Table[Card], Q_SIZE)
	err = DBGetter[Card](db, &cards, DEFAULT_WHERE)
	if err != nil {
		fmt.Println("Error querrying db: ", err)
		return
	}
	col := Col{}
	if err := db.Table("col").Take(&col).Error; err != nil {
		log.Fatal(err)
		return
	}
	today := (int(time.Now().Unix()) - col.Crt) / SECONDS_IN_DAY
	scheduler := InitScheduler()
	IDs, err := scheduler.FillScheduler(&cards, today)
	if err != nil {
		log.Fatal(err)
	}

	flds := make(map[ID]StudyNote, len(cards))
	err = GetFlds(IDs, db, &flds)
	if err != nil {
		return
	}

	scheduler.Study(&cards, db, &flds)
}
