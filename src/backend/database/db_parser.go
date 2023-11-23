package database

import (
	"encoding/json"
	"fmt"
	"log"

	. "src/backend/tables"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gettable interface {
	GetID() ID
}

type Table[T Gettable] map[ID]T

// This function is nearly the same as the one in gorm.sqlite,
// however, it uses a collating function defined in collation/collation.go
func Open(dsn string) gorm.Dialector {
	return &sqlite.Dialector{DSN: dsn, DriverName: "sqlite_unicase"}
}

func OpenDB(filepath string) (db *gorm.DB, err error) {
	db, err = gorm.Open(Open(filepath))
	if err != nil {
		log.Fatal("Error opening DB:", err)
		return
	}
	return
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DBGetter[T Gettable](db *gorm.DB, data *Table[T], table string, where string) error {
	var listData []T
	if where == "" {
		if err := db.Table(table).Find(&listData).Error; err != nil {
			log.Fatal(err)
			return err
		}
	} else {
		if err := db.Table(table).Where(where).Find(&listData).Error; err != nil {
			log.Fatal(err)
			return err
		}
	}

	for i := 0; i < len(listData); i++ {
		(*data)[listData[i].GetID()] = listData[i]
	}
	return nil
}

func ParseDecks(db *gorm.DB, decks *Decks) error {
	cols := Table[Col]{}
	err := DBGetter[Col](db, &cols, "col", "")
	if err != nil {
		log.Fatal(err)
	}
	col := cols[1]
	err = parseJSONDecks(db, &col, decks)
	if err == nil {
		return nil
	}
	return parseDBDecks(db, &col, decks)

}

// unused for now
func parseJSONDecks(db *gorm.DB, col *Col, decks *Decks) error {
	deckStr := col.Decks
	decksJSON := make(map[ID]DeckJSON, 8)
	err := json.Unmarshal([]byte(deckStr), &decksJSON)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return err
	}
	if len(decksJSON) == 0 {
		err := fmt.Errorf("JSON empty")
		log.Println(err)
		return err
	}
	for key := range decksJSON {
		if decksJSON[key].Name == "Default" {
			continue
		}
		deck := Deck{
			ID:   decksJSON[key].ID,
			Name: decksJSON[key].Name,
			Conf: CONFIG,
		}
		GetDeckCardData(&deck, db)
		(*decks)[key] = deck
	}
	return nil
}

func parseDBDecks(db *gorm.DB, col *Col, decks *Decks) error {
	decksTable := Table[DeckTable]{}
	err := DBGetter[DeckTable](db, &decksTable, "decks", "")
	if err != nil {
		log.Println("Error parsing decks:", err)
		return err
	}
	for key := range decksTable {
		if decksTable[key].Name == "Default" {
			continue
		}
		deck := Deck{
			ID:   decksTable[key].ID,
			Name: decksTable[key].Name,
			Conf: CONFIG,
		}
		GetDeckCardData(&deck, db)
		(*decks)[key] = deck
	}
	return err
}

func GetDeckCardData(deck *Deck, db *gorm.DB) error {
	cards := Table[Card]{}
	err := DBGetter[Card](db, &cards, "cards", fmt.Sprintf("did = %d", deck.ID))
	if err != nil {
		log.Println("Error getting cards:", err)
		return err
	}
	deck.New = 0
	deck.Learn = 0
	deck.Due = 0
	for _, card := range cards {
		switch card.Queue {
		case NEW:
			deck.New++
		case LEARNING:
			deck.Learn++
		case REVIEW:
			if TodayRelative(db) < card.Due {
				continue
			}
			deck.Due++
		}
	}
	return nil
}
