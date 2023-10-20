package database

import (
	"database/sql"
	"fmt"
	"os/exec"
)

const TEMPLATE_PATH = "db_media/template.anki2"
const EXPORT_NAME = "export.anki2"

func Export_db(db *sql.DB, decks []int) (err error) {
	err = create_empty_template(db)
	if err != nil {
		fmt.Println("err occured: ", err)
		return
	}
	fmt.Println("template copied")
	return
}

func create_empty_template(db *sql.DB) (err error) {
	cmd := exec.Command("cp", TEMPLATE_PATH, EXPORT_NAME)
	if err = cmd.Run(); err != nil {
		fmt.Println("Error creating template", err)
		return
	}
	return
}
