package sqlite

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type DBSource struct {
	*sql.DB
}

func (b *DBSource) GetDB() *sql.DB {
	return b.DB
}

func New(path string) (*DBSource, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DBSource{db}, nil
}
