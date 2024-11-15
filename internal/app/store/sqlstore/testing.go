package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// TestDB ...
func TestDB(test *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	test.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		test.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		test.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")),
			); err != nil {
				test.Fatal(err)
			}
		}

		if err := db.Close(); err != nil {
			test.Fatal(err)
		}
	}
}
