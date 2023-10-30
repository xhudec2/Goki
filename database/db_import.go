package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mattn/go-sqlite3"
)

const VER = 18

func insert(dest *sql.DB, table string, where_clause string) (err error) {
	_, err = dest.Exec(fmt.Sprintf("INSERT INTO %s SELECT * FROM source_db.%s %s", table, table, where_clause))
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code != 19 || sqliteErr.ExtendedCode != 1555 {
			fmt.Printf("insert %s failed: %v %s", table, err, where_clause)
			return
		}
	}
	return nil
}

func Import_db(dest *sql.DB, source string) (err error) {
	_, err = dest.Exec(fmt.Sprintf("ATTACH DATABASE '%s' AS source_db", source))
	if err != nil {
		fmt.Println("attach failed: ", err)
		return
	}

	if err = insert(dest, "cards", ""); err != nil {
		return
	}

	if err = insert(dest, "notes", ""); err != nil {
		return
	}
	_, err = dest.Exec("DETACH DATABASE source_db")
	return
}

func Create_export_db(dest *sql.DB, source string, decks string) (err error) {
	_, err = dest.Exec(fmt.Sprintf("ATTACH DATABASE '%s' AS source_db", source))
	if err != nil {
		fmt.Println("attach failed: ", err)
		return
	}

	if err = insert(dest, "cards", fmt.Sprintf("WHERE did IN (%s)", decks)); err != nil {
		return
	}

	if err = insert(dest, "decks", fmt.Sprintf("WHERE id IN (%s)", decks)); err != nil {
		return
	}

	if err = insert(dest, "notes", "WHERE id IN (SELECT id FROM cards)"); err != nil {
		return
	}
	curr_time := time.Now().UnixMilli()
	_, err = dest.Exec(fmt.Sprintf("INSERT INTO col VALUES (1, %d, %d, %d, 0, 0, 0, '', '', '', '' , '')", curr_time, curr_time, VER))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dest.Exec("DETACH DATABASE source_db")
	return
}
