package sqlite

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/shared"
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
		return nil, shared.InvalidArgumentException
	}

	uuidUser, err := uuid.NewV7()
	if err != nil {
		return nil, shared.UnknownException
	}

	const (
		intoUserQuery   = `INSERT INTO users(id, name, role) values (?, ?, ?)`
		intoMarkerQuery = `INSERT INTO user_markers(id, user_id, tag) values (?, ?, ?)`
	)

	tx, err := r.DB.Begin()
	if err != nil {
		return nil, shared.InfrastructureException
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
		return nil, shared.InvalidDataException
	}

	u := domain.User{ID: uuidUser.String(), Name: name, Role: 0, Marker: *marker}
	return &u, nil

}

func (r *Repository) UserByMarker(marker *domain.UserMarker) (*domain.User, error) {

	if marker == nil {
		return nil, shared.InvalidArgumentException
	}

	const query = `SELECT u.id, u.name, u.role FROM users u
	INNER JOIN user_markers um ON u.id = um.user_id
	WHERE um.id = ? AND um.tag = ?`

	var user domain.User

	if err := r.GetDB().QueryRow(query, marker.ID, marker.Tag).Scan(&user.ID, &user.Name, &user.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, shared.DoesNotExistsException
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UserByIdString(id string) (*domain.User, error) {
	guid, err := r.parseIdToUUID(id)
	if err != nil {
		return nil, err
	}
	const query = `SELECT u.id, u.name, u.role, um.tag, um.id as marker_id 
	FROM users u
	LEFT JOIN user_markers um ON u.id = um.user_id
    WHERE u.id = ?`

	var user domain.User

	if err := r.GetDB().QueryRow(query, guid[:]).Scan(
		&user.ID,
		&user.Name,
		&user.Role,
		&user.Marker.Tag,
		&user.Marker.ID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, shared.DoesNotExistsException
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) SetUserRoleByIdString(userId string, role domain.Role) error {
	guid, err := r.parseIdToUUID(userId)
	if err != nil {
		return err
	}

	const query = `UPDATE users SET role = ? where id = ?`
	if res, err := r.GetDB().Exec(query, role, guid[:]); err != nil {
		return err
	} else {
		_ = res
	}

	return nil
}

func (r *Repository) parseIdToUUID(id string) (*uuid.UUID, error) {
	if len(id) < 16 {
		return nil, shared.InvalidArgumentException
	}
	guid, err := uuid.Parse(id)
	if err != nil {
		return nil, shared.InvalidArgumentException
	}
	return &guid, nil
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
