package service

import (
	"github.com/assignment/models"
	"gofr.dev/pkg/gofr"
)

type StoreInterface interface {
	GetUserByname(ctx *gofr.Context, name string) (*models.User, error)
	AddUser(ctx *gofr.Context, user *models.User) error
	DeleteUser(ctx *gofr.Context, name string) error
	UpdateUserEmail(ctx *gofr.Context, name string, email string) error
}
