package service

import (
	"Kee-Reall/dnd-manager/internal/domain"
)

type UserService struct {
	containerPtr *Container
}

func (us *UserService) Container() *Container {
	return us.containerPtr
}

func (us *UserService) AccessAndRoleByMarker(m domain.UserMarker) (bool, domain.Role) {
	id64, ok := m.IdInInt64()
	if !ok {
		return false, domain.NoRole
	}

	chatId, ok := us.Container().iVariable("adminChatId")
	if !ok {
		return false, domain.NoRole
	}

	if int64(chatId) == id64 {
		return true, domain.AdminRole
	}
	return false, domain.NoRole
}

func newUserService(cptr *Container) *UserService {
	return &UserService{cptr}
}
