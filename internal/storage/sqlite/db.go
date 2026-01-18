package sqlite

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"database/sql"

	_ "modernc.org/sqlite"
)

type Repository struct {
	*sql.DB
}

func (r *Repository) UserByMarker(marker *domain.UserMarker) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Repository) GetDB() *sql.DB {
	return b.DB
}

func New(path string) (*Repository, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}
