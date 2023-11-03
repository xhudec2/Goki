package database

// https://github.com/ankidroid/Anki-Android/wiki/Database-Structure

// Most of these are unused, only here not to break the code

type Card struct {
	ID     ID `gorm:"primaryKey;autoCreateTime:milli"`
	Nid    int
	Did    int
	Ord    int
	Mod    int `gorm:"autoUpdateTime:milli"`
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

type Col struct {
	ID     ID `gorm:"primaryKey;autoCreateTime:milli"`
	Crt    int
	Mod    int `gorm:"autoUpdateTime:milli"`
	Scm    int
	Ver    int
	Dty    int
	Usn    int
	Ls     int
	Conf   string
	Models string
	Decks  string
	Dconf  string
	Tags   string
}

type Deck struct {
	ID        ID `gorm:"primaryKey;autoCreateTime:milli"`
	Name      string
	MtimeSecs int
	Usn       int
	Common    []byte
	Kind      []byte
}

type Note struct {
	ID    ID `gorm:"primaryKey;autoCreateTime:milli"`
	Guid  string
	Mid   int
	Mod   int `gorm:"autoUpdateTime:milli"`
	Usn   int
	Tags  string
	Flds  string
	Sfld  string
	Csum  int
	Flags int
	Data  string
}

type NoteFlds struct {
	Nid  ID
	Flds string
}

type StudyNote struct {
	Front string
	Back  string
}
