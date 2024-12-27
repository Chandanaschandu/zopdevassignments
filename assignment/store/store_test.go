package store

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/assignment/models"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	"testing"
)

func TestStore_GetUserByname(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	s := New()

	tests := []struct {
		des       string
		name      string
		expOutput *models.User
		expErr    error
	}{
		{
			"Success case",
			"Chandana",
			&models.User{Name: "Chandana", Age: 22, PhoneNumber: "0998898777", Email: "chandana@gmail.com"},
			nil,
		},
		{
			"Failure case (connection failed)",
			"abc",
			&models.User{Name: "", Age: 0, PhoneNumber: "", Email: ""},
			sql.ErrConnDone,
		},
	}
	for i, tc := range tests {
		mock.SQL.ExpectSelect(ctx, tc.expOutput, "SELECT name, age, phoneNumber, email FROM `user` where name=?", tc.name)
		res, err := s.GetUserByname(ctx, tc.name)

		assert.Equalf(t, tc.expErr, err, "Test [%d] Failed: %w", i, tc.des)
		assert.Equal(t, tc.expOutput, res, "Test [%s] Failed:%s", tc.expOutput, res)
	}
}

func Test_Delete(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	s := New()

	tests := []struct {
		desc   string
		name   string
		expErr error
	}{
		{
			"Sucess case", "Chandanna", nil,
		},
		{
			"Failure case: connection closed", "Chandana", sql.ErrConnDone,
		},
	}

	for i, tc := range tests {
		mock.SQL.ExpectExec("DELETE FROM user where name=?").WithArgs(tc.name).WillReturnResult(sqlmock.NewResult(0, 0)).
			WillReturnError(tc.expErr)

		err := s.DeleteUser(ctx, tc.name)

		assert.Equalf(t, tc.expErr, err, "Test [%d] Failed: %s", i, tc.desc)
	}
}

func TestStore_UpdateUserEmail(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	s := New()

	tests := []struct {
		desc   string
		name   string
		email  string
		expErr error
	}{
		{
			"Sucesssfull", "Chandana", "chandana@zop.dev", nil,
		},
		{
			"Failure", "Chandana", "abc", sql.ErrConnDone,
		},
	}
	for i, tt := range tests {
		mock.SQL.ExpectExec("UPDATE `user` SET email=? WHERE name=?").WithArgs(tt.email, tt.name).WillReturnResult(sqlmock.NewResult(0, 0)).
			WillReturnError(tt.expErr)
		err := s.UpdateUserEmail(ctx, tt.name, tt.email)
		assert.Equalf(t, tt.expErr, err, "Test [%d] Failed: %s", i, tt.desc)
	}
}

func TestStore_AddUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
	s := New()
	tests := []struct {
		des    string
		input  *models.User
		expErr error
	}{
		{
			"Success case",
			&models.User{
				Name:        "Chandana",
				Age:         22,
				PhoneNumber: "99877665656",
				Email:       "chanadana@zop.dev",
			}, nil,
		},
		{
			"failure case",
			&models.User{
				Name:        "aaa",
				Age:         99,
				PhoneNumber: "898173781",
				Email:       "aaa@gmail",
			},
			sql.ErrConnDone,
		},
	}
	for i, tt := range tests {
		mock.SQL.ExpectExec("INSERT INTO user(name,age,phoneNumber,email) VALUES (?,?,?,?)").WithArgs(tt.input.Name, tt.input.Age, tt.input.PhoneNumber, tt.input.Email).
			WillReturnResult(sqlmock.NewResult(0, 0)).
			WillReturnError(tt.expErr)

		err := s.AddUser(ctx, tt.input)
		assert.Equalf(t, tt.expErr, err, "Test [%d] Failed: %s", i, tt.des)
	}
}
