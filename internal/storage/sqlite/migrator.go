package sqlite

import (
	"Kee-Reall/dnd-manager/internal/storage"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate(provider storage.DbProvider) error {
	drv, err := sqlite3.WithInstance(provider.GetDB(), &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "", drv)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
