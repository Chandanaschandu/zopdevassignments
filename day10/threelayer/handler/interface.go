package handler

import "github.com/Chandanaschandu/threelayer/models"

type UserServiceInterface interface {
	GetUserByName(name string) (*models.Users, error)
	AddUser(user *models.Users) error
	DeleteUser(name string) error
	UpdateUserEmail(name string, email string) error
}
