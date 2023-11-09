package main

import (
	"fmt"
	"log"
	. "project/collation"
	. "project/database"
	. "project/scheduler"
	. "project/tables"
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
	fmt.Println("Today: ", today)
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
