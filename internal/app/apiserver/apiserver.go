package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Seemann-ng/go-RESTapi/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	store := sqlstore.New(db)
	srv := NewServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

// newDB ...
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
