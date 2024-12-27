package handler

import (
	"github.com/assignment/models"
	"gofr.dev/pkg/gofr"
)

type UserInterface interface {
	GetUserByName(ctx *gofr.Context, name string) (*models.User, error)
	AddUser(ctx *gofr.Context, user *models.User) error
	DeleteUser(ctx *gofr.Context, name string) error
	UpdateEmail(ctx *gofr.Context, name string, email string) error
}
