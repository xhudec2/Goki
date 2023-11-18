package tables

import "log"

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

type DeckTable struct {
	ID         ID `gorm:"primaryKey;autoCreateTime:milli"`
	Name       string
	Mtime_secs int `gorm:"autoUpdateTime"`
	Usn        int `gorm:"default:-1"`
	Common     []byte
	Kind       []byte
}
type Deck struct {
	ID    ID
	Name  string
	Conf  Config
	New   int
	Learn int
	Due   int
}

type Decks map[ID]Deck

func (d DeckTable) GetID() ID {
	return ID(d.ID)
}

func GetDeckID(decks *Decks, deckName string) ID {
	for key := range *decks {
		if (*decks)[key].Name == deckName {
			return key
		}
	}

	log.Fatal("Deck not found")
	return 0
}

func DeckNames(decks *Decks) []string {
	names := make([]string, 0, len(*decks))
	for id := range *decks {
		names = append(names, (*decks)[id].Name)
	}
	return names
}
