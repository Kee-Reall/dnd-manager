package sqlite

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Repository struct {
	*sql.DB
}

func (r *Repository) NewUser(marker *domain.UserMarker, name string) (*domain.User, error) {
	if marker == nil || name == "" {
		return nil, domain.InvalidArgumentException
	}

	uuidUser, err := uuid.NewV7()
	if err != nil {
		return nil, domain.UnknownException
	}

	const (
		intoUserQuery   = `INSERT INTO users(id, name, role) values (?, ?, ?)`
		intoMarkerQuery = `INSERT INTO user_markers(id, user_id, tag) values (?, ?, ?)`
	)

	tx, err := r.DB.Begin()
	if err != nil {
		return nil, domain.InfrastructureException
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.Exec(intoUserQuery, uuidUser[:], name, domain.NoRole); err != nil {
		return nil, err
	}

	if _, err := tx.Exec(intoMarkerQuery, marker.ID, uuidUser[:], marker.Tag); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, domain.InvalidDataException
	}

	u := domain.User{ID: uuidUser.String(), Name: name, Role: 0, Marker: *marker}
	return &u, nil

}

func (r *Repository) UserByMarker(marker *domain.UserMarker) (*domain.User, error) {

	if marker == nil {
		return nil, domain.InvalidArgumentException
	}

	const query = `SELECT u.id, u.name, u.role FROM users u
	INNER JOIN user_markers um ON u.id = um.user_id
	WHERE um.id = ? AND um.tag = ?`

	var user domain.User

	if err := r.GetDB().QueryRow(query, marker.ID, marker.Tag).Scan(&user.ID, &user.Name, &user.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.DoesNotExistsException
		}
		return nil, err
	}

	return &user, nil
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
