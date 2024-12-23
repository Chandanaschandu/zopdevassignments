package service

import (
	"errors"
	"github.com/Chandanaschandu/threelayer/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockStore struct {
}

func TestUserService_GetUser(t *testing.T) {
	mock := mockStore{}
	service := NewServices(mock)

	t.Run("success", func(t *testing.T) {
		user, err := service.GetUserByName("Chandana")

		assert.Nil(t, err)

		expectedUser := &models.Users{
			UserName:    "Chandana",
			UserAge:     12,
			Phonenumber: "9088877766",
			Email:       "chandana@example.com",
		}
		assert.Equal(t, expectedUser, user)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		_, err := service.GetUserByName("unknown")
		assert.ErrorContains(t, err, "user not found")
	})
}

func TestUserService_AddUser(t *testing.T) {
	mock := mockStore{}
	service := NewServices(mock)

	t.Run("success", func(t *testing.T) {
		err := service.AddUser(&models.Users{UserName: "newUser", UserAge: 25, Phonenumber: "9876543210", Email: "new@example.com"})
		assert.Nil(t, err)
	})

	t.Run("failure - duplicate user", func(t *testing.T) {
		err := service.AddUser(&models.Users{UserName: "duplicateUser", UserAge: 25, Phonenumber: "9876543210", Email: "duplicate@example.com"})
		assert.ErrorContains(t, err, "user already exists")
	})
}

func TestUserService_Deleteuser(t *testing.T) {
	mock := mockStore{}
	service := NewServices(mock)

	t.Run("success", func(t *testing.T) {
		err := service.DeleteUser("existingUser")
		assert.Nil(t, err)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		err := service.DeleteUser("nonExistingUser")
		assert.ErrorContains(t, err, "user not found")
	})
}

func TestUserService_UpdateUserEmail(t *testing.T) {
	mock := mockStore{}
	service := NewServices(mock)

	t.Run("success", func(t *testing.T) {
		err := service.UpdateUserEmail("ab", "abcvv@gmail.com")
		assert.Nil(t, err)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		err := service.UpdateUserEmail("nonExistingUser", "abc@gmail.com")
		assert.ErrorContains(t, err, "user not found")
	})
}

func (m mockStore) GetUserName(name string) (*models.Users, error) {
	if name == "Chandana" {
		return &models.Users{UserName: "Chandana", UserAge: 12, Phonenumber: "9088877766", Email: "chandana@example.com"}, nil
	}
	return &models.Users{}, errors.New("user not found")
}

func (m mockStore) AddUser(user *models.Users) error {
	if user.UserName == "duplicateUser" {
		return errors.New("user already exists")
	}
	return nil
}

func (m mockStore) Deleteuser(name string) error {
	if name == "existingUser" {
		return nil
	}
	return errors.New("user not found")
}

func (m mockStore) UpdateUserEmail(name string, email string) error {
	if name == "ab" {
		return nil
	}
	return errors.New("user not found")
}
