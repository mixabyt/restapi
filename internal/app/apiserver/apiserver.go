package apiserver

import (
	"database/sql"
	"mdl/internal/app/store/sqlstore"
	"net/http"

	_ "github.com/lib/pq"
)

func Start(config *Config) error {
	db, err := newDB(config.DataBaseURL)
	if err != nil {
		return err
	}
	store := sqlstore.New(db)
	srv := newServer(store)

	defer db.Close()
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(URL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
