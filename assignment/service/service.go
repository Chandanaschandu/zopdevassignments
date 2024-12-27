package service

import (
	"github.com/assignment/models"
	"gofr.dev/pkg/gofr"
)

type service struct {
	store StoreInterface
}

func New(store StoreInterface) service {
	return service{store: store}
}

func (s service) GetUserByName(ctx *gofr.Context, name string) (*models.User, error) {
	return s.store.GetUserByname(ctx, name)
}

func (s service) AddUser(ctx *gofr.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return s.store.AddUser(ctx, user)
}

func (s service) DeleteUser(ctx *gofr.Context, name string) error {
	return s.store.DeleteUser(ctx, name)
}

func (s service) UpdateEmail(ctx *gofr.Context, name string, email string) error {
	return s.store.UpdateUserEmail(ctx, name, email)
}
