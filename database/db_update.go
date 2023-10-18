package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type Attr_tuple struct {
	Attr_name string
	New_val   string
}

type Updated_attributes []Attr_tuple

func card_to_string(card Card) (str string) {
	v := reflect.ValueOf(card)
	numFields := v.NumField()

	values := make([]string, 0, numFields)

	for i := 0; i < numFields-1; i++ {
		fieldValue := v.Field(i)
		values = append(values, fmt.Sprintf("%v", fieldValue))
	}
	values = append(values, "''")
	return strings.Join(values, ",")
}

func updated_attr_to_string(attrs Updated_attributes) (str string) {
	s := make([]string, 0, 32)
	for _, attr := range attrs {
		s = append(s, fmt.Sprintf("%v = %v", attr.Attr_name, attr.New_val))
	}
	return strings.Join(s, ",")
}

func DB_update(db *sql.DB) {
	// TODO
	// will update the database on startup, check dates and due attrs
}

func Insert_card(db *sql.DB, card Card) (err error) {
	vals := card_to_string(card)
	fmt.Println(vals)
	_, err = db.Exec(fmt.Sprintf("INSERT INTO cards VALUES (%s)", vals))
	if err != nil {
		fmt.Println("Error occured while inserting card: ", err)
		return
	}
	return
}

func Update_card(card_id int, db *sql.DB, attrs Updated_attributes) (err error) {
	_, err = db.Exec(fmt.Sprintf("UPDATE cards SET %s WHERE id = %d", updated_attr_to_string(attrs), card_id))
	if err != nil {
		fmt.Printf("Error updating card %d, err: %v", card_id, err)
		return
	}
	return
}

func Delete_card(card_id int, db *sql.DB) (err error) {
	rows, err := db.Query(fmt.Sprintf("SELECT count(*), did, nid FROM cards WHERE id = %d GROUP BY did", card_id))
	if err != nil {
		fmt.Println("Error querrying db: ", err)
		return
	}
	defer rows.Close()
	var count, did, nid int
	rows.Next()
	err = rows.Scan(&count, &did, &nid)

	if err != nil {
		fmt.Println("Error scanning db: ", err)
		return
	}

	if count == 1 && (Delete_note(nid, db) != nil) {
		return
	}

	_, err = db.Exec(fmt.Sprintf("DELETE FROM cards WHERE id = %d", card_id))
	if err != nil {
		fmt.Println("Error deleting card: ", err)
		return
	}
	fmt.Println("deleted card ", card_id)
	return
}

func Delete_note(note_id int, db *sql.DB) (err error) {
	_, err = db.Exec(fmt.Sprintf("DELETE FROM notes WHERE id = %d", note_id))
	if err != nil {
		fmt.Println("Error deleting card: ", err)
		return
	}
	fmt.Println("deleting note ", note_id)
	return
}
