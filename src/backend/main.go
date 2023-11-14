package main

import (
	"fmt"
	"log"
	. "src/backend/collation"
	. "src/backend/database"
	. "src/backend/scheduler"
	. "src/backend/tables"
)

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
	today := TodayRelative(db)
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
	scheduler.Study(&cards, db, &CONFIG, &flds)
}
