package storage

import "database/sql"

type DbProvider interface {
	GetDB() *sql.DB
}
