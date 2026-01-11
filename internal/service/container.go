package service

import "database/sql"

type ContainerProvider interface {
	Container() *Container
}

type Container struct {
	userService *UserService
	sVars       map[string]string
	intVars     map[string]int
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

type DBProvider interface {
	GetDB() *sql.DB
}

func NewContainer(db DBProvider, admChatId int64) (*Container, error) {
	sVars, intVars := make(map[string]string), make(map[string]int, 1)
	intVars[adminChatId] = int(admChatId)
	c := &Container{
		sVars:   sVars,
		intVars: intVars,
	}
	c.userService = newUserService(c)
	return c, nil
}
