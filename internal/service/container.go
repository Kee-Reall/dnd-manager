package service

import (
	"Kee-Reall/dnd-manager/internal/config"
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
	boolVars    map[string]bool
	repo        *sqlite.Repository
}

func (c *Container) IVariable(vName string) (int, bool) {
	v, ok := c.intVars[vName]
	return v, ok
}

func (c *Container) SVariable(vName string) (string, bool) {
	v, ok := c.sVars[vName]
	return v, ok
}

func (c *Container) BVariable(vName string) (bool, bool) {
	v, ok := c.boolVars[vName]
	return v, ok
}

func (c *Container) UserService() *UserService {
	return c.userService
}

func NewContainer(repo *sqlite.Repository, cfg *config.Config) (*Container, error) {
	sVars, intVars, bVars := make(map[string]string), make(map[string]int, 1), make(map[string]bool, 1)
	intVars[AdminChatId] = int(cfg.AdminChatId)
	bVars[RegEnable] = cfg.RegEnable
	c := &Container{
		sVars:    sVars,
		intVars:  intVars,
		boolVars: bVars,
		repo:     repo,
	}
	c.userService = newUserService(c)
	return c, nil
}
