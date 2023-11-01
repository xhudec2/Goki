package database

type Card struct {
	ID     uint `gorm:"primaryKey"`
	Nid    int
	Did    int
	Ord    int
	Mod    int
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
	ID     int `gorm:"primaryKey"`
	Crt    int
	Mod    int
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
	ID        int `gorm:"primaryKey"`
	Name      string
	MtimeSecs int
	Usn       int
	Common    []byte
	Kind      []byte
}

type Note struct {
	ID    int `gorm:"primaryKey"`
	Guid  string
	Mid   int
	Mod   int
	Usn   int
	Tags  string
	Flds  string
	Sfld  string
	Csum  int
	Flags int
	Data  string
}
