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
