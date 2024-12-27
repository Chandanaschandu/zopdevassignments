package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/assignment/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"testing"
)

func TestUserService_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStoreInterface(ctrl)
	service := New(mockStore)

	tests := []struct {
		name          string
		mockSetup     func()
		user          *models.User
		expectedError error
	}{
		{
			name: "User Added Successfully",
			mockSetup: func() {
				mockStore.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			user: &models.User{
				Name:        "john",
				Age:         30,
				PhoneNumber: "1234567890",
				Email:       "john@example.com",
			},
			expectedError: nil,
		},
		{
			name: "User Validation Failed",
			mockSetup: func() {
			},
			user: &models.User{
				Name:        "11",
				Age:         30,
				PhoneNumber: "1234567",
				Email:       "john@example.com",
			},
			expectedError: errors.New("validation failed: name cannot be empty"),
		},
		{
			name: "Store Error on Add",
			mockSetup: func() {
				mockStore.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			user: &models.User{
				Name:        "john",
				Age:         30,
				PhoneNumber: "1234567890",
				Email:       "john@example.com",
			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup()
			ctx := &gofr.Context{
				Context: context.Background(),
			}

			err := service.AddUser(ctx, tt.user)

			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestUserService_GetUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStoreInterface(ctrl)
	service := New(mockStore)
	tests := []struct {
		name          string
		mockSetup     func()
		userName      string
		expectedUser  *models.User
		expectedError error
	}{
		{
			name: "User Found",
			mockSetup: func() {
				mockStore.EXPECT().GetUserByname(gomock.Any(), "john").Return(&models.User{
					Name:        "john",
					Age:         30,
					PhoneNumber: "1234567890",
					Email:       "john@example.com",
				}, nil)
			},
			userName: "john",
			expectedUser: &models.User{
				Name:        "john",
				Age:         30,
				PhoneNumber: "1234567890",
				Email:       "john@example.com",
			},
			expectedError: nil,
		},
		{
			name: "User Not Found",
			mockSetup: func() {
				mockStore.EXPECT().GetUserByname(gomock.Any(), "john").Return(nil, sql.ErrNoRows)
			},
			userName:      "john",
			expectedUser:  nil,
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := &gofr.Context{
				Context: context.Background(),
			}

			result, err := service.GetUserByName(ctx, tt.userName)

			assert.Equal(t, tt.expectedUser, result)

			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStoreInterface(ctrl)
	service := New(mockStore)
	tests := []struct {
		name          string
		mockSetup     func()
		userName      string
		expectedError error
	}{
		{
			name: "User Successfully Deleted",
			mockSetup: func() {
				mockStore.EXPECT().DeleteUser(gomock.Any(), "john").Return(nil)
			},
			userName:      "john",
			expectedError: nil,
		},
		{
			name: "Store Error",
			mockSetup: func() {
				mockStore.EXPECT().DeleteUser(gomock.Any(), "john").Return(errors.New("database error"))
			},
			userName:      "john",
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := &gofr.Context{
				Context: context.Background(),
			}

			err := service.DeleteUser(ctx, tt.userName)
			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestUserService_UpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStoreInterface(ctrl)
	service := New(mockStore)

	tests := []struct {
		name          string
		mockSetup     func()
		userName      string
		email         string
		expectedError error
	}{
		{
			name: "Email Successfully Updated",
			mockSetup: func() {
				mockStore.EXPECT().UpdateUserEmail(gomock.Any(), "abc", "abc@example.com").Return(nil)
			},
			userName:      "abc",
			email:         "abc@example.com",
			expectedError: nil,
		},
		{
			name: "Store Error",
			mockSetup: func() {
				mockStore.EXPECT().UpdateUserEmail(gomock.Any(), "john", "john_new@example.com").Return(errors.New("database error"))
			},
			userName:      "john",
			email:         "john_new@example.com",
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := &gofr.Context{
				Context: context.Background(),
			}

			err := service.UpdateEmail(ctx, tt.userName, tt.email)

			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			}
		})
	}
}
