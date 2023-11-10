package tables

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

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
	NEW            = 0
	LEARNING       = 1
	REVIEW         = 2
	USER_SUSPENDED = 3
	SUSPENDED      = -1
)
const (
	DAY    = 86400
	MINUTE = 60
	HOUR   = 3600
)

const CARD_DELIMITER = "\x1f"

func StudyCard(card *Card, db *gorm.DB, conf *Config, flds *map[ID]StudyNote) (bool, error) {
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
	grade, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return false, err
	}
	// updateCard

	return card.UpdateCard(grade, db, conf), nil
}
