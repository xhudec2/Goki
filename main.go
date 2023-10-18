package main

import (
	"project/database"
)

func main() {

	db, err := database.Open_db("Anki2/User 1/collection.anki2")
	if err != nil {
		return
	}
	defer db.Close()

}
