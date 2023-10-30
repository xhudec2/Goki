package database

import "strings"

func unicaseCollation(a, b string) int {
	// Your custom collation logic (case-insensitive comparison)
	return strings.Compare(strings.ToLower(a), strings.ToLower(b))
}
