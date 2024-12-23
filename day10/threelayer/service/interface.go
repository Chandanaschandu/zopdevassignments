package service

import "github.com/Chandanaschandu/threelayer/models"

type StoreInterface interface {
	GetUserName(name string) (*models.Users, error)
	AddUser(user *models.Users) error
	Deleteuser(name string) error
	UpdateUserEmail(name, email string) error
}
