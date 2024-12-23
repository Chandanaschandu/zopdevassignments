package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Chandanaschandu/threelayer/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
}

func TestUserHandler_GetUserByName(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/user/ab", nil)
		req = mux.SetURLVars(req, map[string]string{"user_name": "ab"})

		rec := httptest.NewRecorder()

		handler.GetUserByName(rec, req)

		assert.Equal(t, rec.Code, http.StatusOK)

		expectedBody := `{"user_name":"ab","user_age":31,"phone_Number":"8090892381","email":"abcvv@gmail.com"}`
		assert.JSONEq(t, expectedBody, rec.Body.String())
	})

	t.Run("failure", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/user/invalidUser", nil)
		req = mux.SetURLVars(req, map[string]string{"user_name": "invalidUser"})

		rec := httptest.NewRecorder()

		handler.GetUserByName(rec, req)

		assert.Equal(t, rec.Code, http.StatusNotFound)

		expectedBody := `{"error":"User not found"}`
		assert.JSONEq(t, expectedBody, rec.Body.String())
	})
}

func TestUserHandler_AddUser(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		user := models.Users{UserName: "newUser", UserAge: 30, Phonenumber: "9876543210", Email: "new@gmail.com"}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		handler.AddUsers(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
		expectedBody := `{"user_name":"newUser","user_age":30,"phone_Number":"9876543210","email":"new@gmail.com"}
`
		assert.Equal(t, rec.Body.String(), expectedBody)
	})

	t.Run("failure - duplicate user", func(t *testing.T) {
		user := models.Users{UserName: "duplicateUser", UserAge: 30, Phonenumber: "9876543210", Email: "duplicate@example.com"}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		handler.AddUsers(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("failure - invalid body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte("{invalid json}")))
		rec := httptest.NewRecorder()
		handler.AddUsers(rec, req)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})
}

func TestUserHandler_DeleteUser(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/user/validUser", nil)
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.DeleteUsers(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/user/invalidUser", nil)
		req = mux.SetURLVars(req, map[string]string{"name": "invalidUser"})
		rec := httptest.NewRecorder()
		handler.GetUserByName(rec, req)
		assert.Equal(t, rec.Code, http.StatusNotFound)
	})
}

func TestUserHandler_UpdateUserEmail(t *testing.T) {
	handler := NewUserHandler(mockService{})

	t.Run("success", func(t *testing.T) {
		user := models.Users{UserName: "updatedUser", UserAge: 35, Phonenumber: "1231231234", Email: "updated@example.com"}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPut, "/user/validUser", bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.UpdateUserEmail(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
	})

	t.Run("failure - user not found", func(t *testing.T) {
		user := models.Users{UserName: "unknownUser", UserAge: 35, Phonenumber: "1231231234", Email: "unknown@example.com"}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPut, "/user/invalidUser", bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"name": "invalidUser"})
		rec := httptest.NewRecorder()
		handler.GetUserByName(rec, req)
		assert.Equal(t, rec.Code, http.StatusNotFound)
	})

	t.Run("failure - invalid body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/user/validUser", bytes.NewBuffer([]byte("{invalid json}")))
		req = mux.SetURLVars(req, map[string]string{"name": "validUser"})
		rec := httptest.NewRecorder()
		handler.GetUserByName(rec, req)
		assert.Equal(t, rec.Code, http.StatusNotFound)
	})
}

func (m mockService) GetUserByName(name string) (*models.Users, error) {
	if name == "ab" {

		return &models.Users{
			UserName:    "ab",
			UserAge:     31,
			Phonenumber: "8090892381",
			Email:       "abcvv@gmail.com",
		}, nil
	}

	return &models.Users{}, fmt.Errorf("user not found")
}

func (m mockService) AddUser(user *models.Users) error {
	if user.UserName == "duplicateUser" {
		return errors.New("user already exists")
	}
	return nil
}

func (m mockService) DeleteUser(name string) error {
	if name != "validUser" {
		return errors.New("user not found")
	}
	return nil
}

func (m mockService) UpdateUserEmail(name string, email string) error {
	if name != "validUser" {
		return errors.New("user not found")
	}
	return nil
}
