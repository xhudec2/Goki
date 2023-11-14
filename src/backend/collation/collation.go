package collation

import (
	"database/sql"
	"strings"

	"github.com/mattn/go-sqlite3"
)

func comparator(a, b string) int {
	return strings.Compare(strings.ToLower(a), strings.ToLower(b))
}

func hook(conn *sqlite3.SQLiteConn) error {
	return conn.RegisterCollation("unicase", comparator)
}

// Needed for the database to work properly.
func RegisterCollation() {
	sql.Register(
		"sqlite_unicase",
		&sqlite3.SQLiteDriver{ConnectHook: hook})
}
