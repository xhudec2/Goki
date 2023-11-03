package main

import (
	"fmt"
	"log"
	. "project/collation"
	. "project/database"
	. "project/scheduler"
)

func main() {
	RegisterCollation()
	db, err := OpenDB("export/1")
	if err != nil {
		return
	}
	defer CloseDB(db)

	//	export, err := OpenDB("test/export.anki2")
	//	if err != nil {
	//		return
	//	}
	//	ImportDB[Card](db, export, "cards")
	// ExportDB("test", "test")

	cards := make(Table[Card], Q_SIZE)
	err = DBGetter[Card](db, &cards, DEFAULT_WHERE)
	if err != nil {
		fmt.Println("Error querrying db: ", err)
		return
	}
	scheduler := SchedulerInit()
	err = scheduler.FillScheduler(&cards)
	if err != nil {
		log.Fatal(err)
	}
	scheduler.Study(&cards, db)
}
