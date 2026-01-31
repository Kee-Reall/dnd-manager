package service

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"errors"
)

type UserService struct {
	containerPtr *Container
}

func (us *UserService) Container() *Container {
	return us.containerPtr
}

func (us *UserService) RegisterNewUserByTag(tgId, name string) error {
	um := &domain.UserMarker{tgId, "tg"}
	dbUser, err := us.Container().repo.UserByMarker(um)

	var user *domain.User = nil
	switch {
	case err == nil:
		if dbUser.Role == domain.NoRole {
			return domain.NotAllowedException
		}
		return domain.UnknownException
	case errors.Is(err, domain.DoesNotExistsException):
		if user, err = us.Container().repo.NewUser(um, name); err != nil {
			return err
		}
	default:
		return err
	}

	us.Container().EventBus.Publish(*user, "new-user-reg")

	return nil

}

func (us *UserService) UserByMarker(m *domain.UserMarker) (*domain.User, error) {
	if m == nil {
		return nil, domain.InvalidArgumentException
	}

	return us.Container().repo.UserByMarker(m)
}

func (us *UserService) AccessAndRoleByMarker(m domain.UserMarker) (bool, domain.Role) {
	id64, ok := m.IdInInt64()
	if !ok {
		return false, domain.NoRole
	}

	chatId, ok := us.Container().IVariable("adminChatId")
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
