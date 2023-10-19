package database

import (
	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

func insert(main_db *sql.DB, table string) (err error) {
	_, err = main_db.Exec(fmt.Sprintf("INSERT INTO %s SELECT * FROM source_db.%s", table, table))
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code != 19 || sqliteErr.ExtendedCode != 1555 {
			fmt.Printf("insert %s failed: %v", table, err)
			return
		}
	}
	return nil
}

func Import_db(main_db *sql.DB, source string) (err error) {
	_, err = main_db.Exec(fmt.Sprintf("ATTACH DATABASE '%s' AS source_db", source))
	if err != nil {
		fmt.Println("attach failed: ", err)
		return
	}

	if err = insert(main_db, "cards"); err != nil {
		return
	}

	if err = insert(main_db, "notes"); err != nil {
		return
	}
	_, err = main_db.Exec("DETACH DATABASE source_db")
	return
}
