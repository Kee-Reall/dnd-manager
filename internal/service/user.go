package service

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/shared"
	"Kee-Reall/dnd-manager/internal/storage/sqlite"
	"errors"
	"strconv"
)

type Publisher interface {
	Publish(any, string)
}

type UserService struct {
	containerPtr *Container
	repo         *sqlite.Repository
	bus          Publisher
}

func (us *UserService) Container() *Container {
	return us.containerPtr
}

func (us *UserService) RegisterNewUserByTag(tgId, name string) error {
	um := &domain.UserMarker{tgId, "tg"}
	dbUser, err := us.repo.UserByMarker(um)

	var user *domain.User = nil
	switch {
	case err == nil:
		if dbUser.Role == domain.NoRole {
			return shared.NotAllowedException
		}
		return shared.UnknownException
	case errors.Is(err, shared.DoesNotExistsException):
		if user, err = us.repo.NewUser(um, name); err != nil {
			return err
		}
	default:
		return err
	}

	us.Container().EventBus.Publish(*user, "new-user-reg")

	return nil

}

func (us *UserService) AcceptUser(id string) error {
	u, err := us.repo.UserByIdString(id)
	if err != nil {
		return shared.InvalidDataException
	}
	if u.Role != domain.NoRole {
		if u.Role == domain.PlayerRole {
			return shared.ScenarioAlreadyDoneException
		}
		return shared.NotAllowedException //уже подвержден
	}
	if err := us.repo.SetUserRoleByIdString(id, domain.PlayerRole); err != nil {
		return err
	}

	us.Container().EventBus.Publish(*u, "user-accepted")

	return nil
}

func (us *UserService) UserByMarker(m *domain.UserMarker) (*domain.User, error) {
	if m == nil {
		return nil, shared.InvalidArgumentException
	}

	return us.repo.UserByMarker(m)
}

func (us *UserService) AccessAndRoleByMarker(m *domain.UserMarker) (bool, *domain.User) {
	id64, ok := m.IdInInt64()
	if !ok {
		return false, nil
	}

	adminChatId, ok := us.Container().IVariable("adminChatId")
	if !ok {
		return false, nil
	}

	if int64(adminChatId) == id64 {
		return true, &domain.User{
			ID:     "0",
			Name:   "GOD",
			Role:   domain.AdminRole,
			Marker: domain.UserMarker{strconv.Itoa(adminChatId), "tg"},
		}
	}

	user, err := us.repo.UserByMarker(m)
	if err != nil {
		return false, nil
	}

	if user != nil && user.Role != domain.AdminRole {
		return true, user
	}

	return false, nil
}

func newUserService(cptr *Container) *UserService {
	return &UserService{cptr, cptr.repo, cptr.EventBus}
}
