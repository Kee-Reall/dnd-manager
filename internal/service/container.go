package service

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/storage/sqlite"
)

type ContainerProvider interface {
	Container() *Container
}

type UserByMarkerProvider interface {
	UserByMarker(marker *domain.UserMarker) (*domain.User, error)
}

type Container struct {
	userService *UserService
	sVars       map[string]string
	intVars     map[string]int
	repo        *sqlite.Repository
}

func (c *Container) iVariable(vName string) (int, bool) {
	v, ok := c.intVars[vName]
	return v, ok
}

func (c *Container) sVariable(vName string) (string, bool) {
	v, ok := c.sVars[vName]
	return v, ok
}

func (c *Container) UserService() *UserService {
	return c.userService
}

func NewContainer(repo *sqlite.Repository, admChatId int64) (*Container, error) {
	sVars, intVars := make(map[string]string), make(map[string]int, 1)
	intVars[adminChatId] = int(admChatId)
	c := &Container{
		sVars:   sVars,
		intVars: intVars,
		repo:    repo,
	}
	c.userService = newUserService(c)
	return c, nil
}
