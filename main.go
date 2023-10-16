package main

import (
	. "project/collation"
	. "project/database"
)

func main() {
	RegisterCollation()
	//	db, err := OpenDB(MAIN_DB)
	//	if err != nil {
	//		return
	//	}
	//	export, err := OpenDB("test/export.anki2")
	//	if err != nil {
	//		return
	//	}
	//	ImportDB[Card](db, export, "cards")
	ExportDB("test", "test")
}

func main() {
	//test_db()
	db, err := database.Open_db("Anki2/User 1/collection.anki2")
	if err != nil {
		return
	}
	defer db.Close()
	qs := scheduler.Scheduler_init()
	err = qs.Fill_scheduler(db)
	if err != nil {
		return
	}
	err = qs.Study(db)
	if err != nil {
		return
	}
}
