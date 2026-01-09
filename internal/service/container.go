package service

import (
	"database/sql"
)

type Container struct{}

type DBProvider interface {
	GetDB() *sql.DB
}

func NewContainer(db DBProvider) (*Container, error) {
	return &Container{}, nil
}
