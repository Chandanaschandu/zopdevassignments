package store

import (
	"github.com/assignment/models"

	"gofr.dev/pkg/gofr"
)

type store struct {
}

func New() *store {
	return &store{}
}

func (s *store) GetUserByname(ctx *gofr.Context, name string) (*models.User, error) {
	var user models.User

	ctx.SQL.Select(ctx, &user, "SELECT name, age, phoneNumber, email FROM `user` where name=?", name)

	return &user, nil

}

func (s *store) AddUser(ctx *gofr.Context, user *models.User) error {
	_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO user(name,age,phoneNumber,email) VALUES (?,?,?,?)", user.Name, user.Age, user.PhoneNumber, user.Email)
	return err
}

func (s *store) DeleteUser(ctx *gofr.Context, name string) error {
	_, err := ctx.SQL.ExecContext(ctx, "DELETE FROM user where name=?", name)
	return err
}

func (s *store) UpdateUserEmail(ctx *gofr.Context, name string, email string) error {

	_, err := ctx.SQL.ExecContext(ctx, "UPDATE `user` SET email=? WHERE name=?", email, name)
	return err

}
