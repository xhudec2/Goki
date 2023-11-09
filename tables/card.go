package tables

import (
	"bufio"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type Card struct {
	ID     ID `gorm:"primaryKey;autoCreateTime:milli"`
	Nid    ID
	Did    ID
	Ord    int
	Mod    int `gorm:"autoUpdateTime"`
	Usn    int
	Type   int
	Queue  int
	Due    int
	Ivl    int
	Factor int
	Reps   int
	Lapses int
	Left   int
	Odue   int
	Odid   int
	Flags  int
	Data   string
}

func (c Card) GetID() ID {
	return ID(c.ID)
}

const (
	NEW       = 0
	LEARNING  = 1
	REVIEW    = 2
	SUSPENDED = -1
	MINUTE    = 60
)

const CARD_DELIMITER = "\x1f"

func StudyCard(card *Card, db *gorm.DB, flds *map[ID]StudyNote) (bool, error) {
	// TODO: unfinished, only prints the card for now
	// It should also change the db based on the "grade" given to the card while studying
	// return the grade as well to know where to store the card next

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(">", (*flds)[card.ID].Front)
	scanner.Scan()
	_ = scanner.Text() //empty line

	fmt.Println((*flds)[card.ID].Back)
	fmt.Print("Grade: ")
	scanner.Scan()
	grade := scanner.Text()

	// updateCard

	return card.UpdateCard(db, grade), nil
}
