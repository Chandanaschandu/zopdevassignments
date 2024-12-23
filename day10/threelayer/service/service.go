package service

import (
	"github.com/Chandanaschandu/threelayer/models"
)

type userService struct {
	store StoreInterface
}

func NewServices(store StoreInterface) userService {
	return userService{store: store}
}

func (s userService) GetUserByName(name string) (*models.Users, error) {
	return s.store.GetUserName(name)
}

func (s userService) AddUser(user *models.Users) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return s.store.AddUser(user)
}

func (s userService) DeleteUser(name string) error {
	return s.store.Deleteuser(name)
}

func (s userService) UpdateUserEmail(name, email string) error {
	return s.store.UpdateUserEmail(name, email)
}
