package tables

import (
	"encoding/json"
	"fmt"
)

type newConf struct {
	Delays []int
	Ints   []int
	Factor int
}

type reviewStruct struct {
	HardFactor float64
	EasyFactor float64
	EasyBonus  float64
}

type lapseStruct struct {
	Delays     []int
	LeechLimit int
	Mult       float64
}

type Config struct {
	New    newConf
	Review reviewStruct
	Lapse  lapseStruct
}

var CONFIG = Config{
	New: newConf{
		// user can set delays and ints to be whatever they want
		// delays can be however long
		Delays: []int{1, 10},
		Ints:   []int{1, 4},
		Factor: 2500,
	},
	Review: reviewStruct{
		HardFactor: 1.2,
		EasyBonus:  1.3,
	},
	Lapse: lapseStruct{
		Delays:     []int{10},
		LeechLimit: 8,
		Mult:       0,
	},
}

// newer versions of the database do not have a table called decks,
// instead they store this data in Col table where decks is a JSON string...
type DeckJSON struct {
	ID               ID
	Mod              int
	Name             string
	Usn              int
	LrnToday         []int
	RevToday         []int
	NewToday         []int
	TimeToday        []int
	Collapsed        bool
	BrowserCollapsed bool
	Desc             string
	Dyn              int
	Conf             int
	ExtendNew        int
	ExtendRev        int
	ReviewLimit      int
	NewLimit         int
	ReviewLimitToday int
	NewLimitToday    int
}

type Deck struct {
	ID   ID
	Name string
	Conf Config
}

type Decks map[ID]Deck

func (d Deck) GetID() ID {
	return ID(d.ID)
}

// unused for now
func ParseDecks(col *Col, decks *Decks) {
	deckStr := col.Decks
	decksJSON := make(map[ID]DeckJSON, 8)
	err := json.Unmarshal([]byte(deckStr), &decksJSON)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	for key := range decksJSON {
		(*decks)[key] = Deck{
			ID:   decksJSON[key].ID,
			Name: decksJSON[key].Name,
			Conf: CONFIG, // default config only for now
		}
	}
}
