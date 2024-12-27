package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/assignment/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	gofrhttp "gofr.dev/pkg/gofr/http"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockUserInterface(ctrl)
	h := New(mockService)

	tests := []struct {
		name             string
		body             interface{}
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			name: "successful user addition",
			body: &models.User{Name: "sbb", Age: 30, PhoneNumber: "123456789", Email: "sbb@example.com"},
			mockExpect: func() {
				mockService.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			expectedResponse: nil,
			expectedErr:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte(`{"name":"sbb","age":30,"phone":"123456789","email":"sbb@example.com"}`)))
			req.Header.Set("Content-Type", "application/json")

			c := &gofr.Context{
				Context: nil,
				Request: gofrhttp.NewRequest(req),
			}

			if tt.body != nil {
				c.Bind(tt.body)
			}

			tt.mockExpect()

			res, err := h.AddUser(c)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResponse, res)
		})
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserInterface(ctrl)
	handler := New(mock)

	tests := []struct {
		des      string
		mockfunc func()
		path     string
		err      error
	}{
		{
			"Success",
			func() {
				mock.EXPECT().DeleteUser(gomock.Any(), "Chandana").Return(nil)
			},
			"Chandana",
			nil,
		},
		{
			"Failure case",
			func() {
				mock.EXPECT().DeleteUser(gomock.Any(), "ab").Return(gofrhttp.ErrorEntityNotFound{})
			},
			"ab",
			gofrhttp.ErrorEntityNotFound{},
		},
		{
			"empty path",
			func() {
				mock.EXPECT().DeleteUser(gomock.Any(), "").Return(gofrhttp.ErrorMissingParam{})
			},
			"",
			gofrhttp.ErrorMissingParam{},
		},
	}
	for _, tc := range tests {

		t.Run(tc.des, func(t *testing.T) {
			tc.mockfunc()
			req := httptest.NewRequest(http.MethodDelete, "/users/"+tc.path, http.NoBody)
			gofrR := gofrhttp.NewRequest(gofrhttp.SetPathParam(req, map[string]string{"name": tc.path}))
			ctx := &gofr.Context{
				Context: context.Background(),
				Request: gofrR,
			}
			tc.mockfunc()
			_, err := handler.DeleteUser(ctx)

			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tc.err.Error(), err.Error())
			}

		})
	}
}

func TestHandler_GetUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserInterface(ctrl)
	handler := New(mock)

	tests := []struct {
		des              string
		pathParam        string
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{"Success", "Chandana",
			func() {
				mock.EXPECT().GetUserByName(gomock.Any(), "Chandana").
					Return(&models.User{
						Name:        "Chandana",
						Age:         22,
						PhoneNumber: "9987898744",
						Email:       "chandu@gmail.com",
					}, nil)
			},
			&models.User{
				Name:        "Chandana",
				Age:         22,
				PhoneNumber: "9987898744",
				Email:       "chandu@gmail.com",
			}, nil,
		},
		{

			"Failure", "ab",
			func() {
				mock.EXPECT().GetUserByName(gomock.Any(), "ab").Return(nil, errors.New("user not found"))
			},
			nil,
			errors.New("user not found"),
		},
		{
			"failure", "",
			func() {
				mock.EXPECT().GetUserByName(gomock.Any(), "").Return(nil, gofrhttp.ErrorMissingParam{})
			},
			nil,
			gofrhttp.ErrorMissingParam{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.des, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/users/"+tt.pathParam, nil)
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrhttp.NewRequest(gofrhttp.SetPathParam(req, map[string]string{"name": tt.pathParam}))

			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}
			tt.mockExpect()

			res, err := handler.GetUserByName(c)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResponse, res)
		})
	}

}

func TestHandler_UpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserInterface(ctrl)
	handler := New(mock)

	tests := []struct {
		des              string
		pathParam        string
		body             interface{}
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			"Success case",
			"Chandana",
			&models.User{
				Name:        "Chandana",
				Age:         22,
				PhoneNumber: "89877666532",
				Email:       "chandana@gmail.com",
			},
			func() {
				mock.EXPECT().UpdateEmail(gomock.Any(), "Chandana", "chandana@gmail.com").Return(nil)
			},
			nil,
			nil,
		},
		{
			"Failure case",
			"aa",
			&models.User{Name: "aa", Age: 2, PhoneNumber: "00892222211", Email: ""},
			func() {
				mock.EXPECT().UpdateEmail(gomock.Any(), "aa", "").Return(gofrhttp.ErrorInvalidParam{})
			},
			nil,
			gofrhttp.ErrorInvalidParam{},
		},
		{
			"Failure case 2",
			"/abc",
			&models.User{Name: "abc", Age: 22, PhoneNumber: "9992222211", Email: "cc@gmail.com"},
			func() {
				mock.EXPECT().UpdateEmail(gomock.Any(), "/abc", "cc@gmail.com").Return(gofrhttp.ErrorInvalidParam{})
			},
			nil,
			gofrhttp.ErrorInvalidParam{},
		},
		//{
		//	"Failure in reading the body",
		//	"chandu",
		//	&models.User{
		//		Name: "chandu", Age: 0, PhoneNumber: "9229292922", Email: "123",
		//	},
		//	func() {
		//		mock.EXPECT().UpdateEmail(gomock.Any(), "chandu", "123").Return(gofrhttp.ErrorInvalidParam{})
		//	},
		//	nil,
		//	gofrhttp.ErrorInvalidParam{},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.des, func(t *testing.T) {

			var req *http.Request
			if tt.body != nil {
				body, err := json.Marshal(tt.body)
				assert.NoError(t, err)
				req = httptest.NewRequest(http.MethodPut, "/user/"+tt.pathParam, bytes.NewReader(body))
			} else {
				req = httptest.NewRequest(http.MethodPut, "/user/"+tt.pathParam, http.NoBody)
			}
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrhttp.NewRequest(gofrhttp.SetPathParam(req, map[string]string{"name": tt.pathParam}))

			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}

			tt.mockExpect()

			res, err := handler.UpdateEmail(c)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResponse, res)
		})
	}
}
