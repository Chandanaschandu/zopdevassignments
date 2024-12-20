package service

import (
	"github.com/Chandanaschandu/threelayer/models"
	"github.com/Chandanaschandu/threelayer/store"
)

type UserMethods interface {
	GetUserByName(name string) (models.Users, error)
	AddUser(user models.Users) error
	DeleteUser(name string) error
	UpdateUserEmail(name string, email string) error
}

type UserService struct {
	store *store.Store
}

func NewServices(store *store.Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) GetUserByName(name string) (*models.Users, error) {
	return s.store.GetUserName(name)
}

func (s *UserService) AddUsers(user *models.Users) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return s.store.AddUser(user)
}

func (s *UserService) DeleteUser(name string) error {
	return s.store.Deleteuser(name)
}

func (s *UserService) UpdateUserEmail(name, user string) error {
	return s.store.UpdateUserEmail(name, user)
}
