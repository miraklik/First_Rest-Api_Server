package apiserver

import (
	"database/sql"
	"golangApi/https-rest-api/internal/app/store/sqlstore"
	"net/http"

	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	db, err := newDB(config.DataBase)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessionStore)

	return http.ListenAndServe(config.BinAddr, srv)
}

func NewDB(dataBaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataBaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
