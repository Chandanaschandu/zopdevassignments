package store

import (
	"fmt"
	"github.com/Chandanaschandu/threelayer/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestUserStore_GetUserName(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error in sql mock %v", err)
	}
	defer db.Close()

	store := NewUserStore(db)

	t.Run("success", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"UserName", "UserAge", "Phone_number", "Email"}).
			AddRow("Chandana", 20, "9019878823", "chandana@gmail.com")
		mock.ExpectQuery("SELECT UserName, UserAge, Phone_number, Email FROM User WHERE UserName = ?").
			WithArgs("Chandana").
			WillReturnRows(rows)

		user, err := store.GetUserName("Chandana")

		assert.NoError(t, err)
		expectedUser := &models.Users{
			UserName:    "Chandana",
			UserAge:     20,
			Phonenumber: "9019878823",
			Email:       "chandana@gmail.com",
		}
		assert.Equal(t, expectedUser, user)

	})

	t.Run("failure", func(t *testing.T) {

		mock.ExpectQuery("SELECT UserName, UserAge, Phone_number, Email FROM User WHERE UserName = ?").
			WithArgs("Chandana").
			WillReturnError(fmt.Errorf("error in database"))

		user, err := store.GetUserName("Chandana")

		assert.Error(t, err)
		assert.Nil(t, user)

		assert.EqualError(t, err, "error fetching user: error in database")

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func Test_AddUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error in sql mock %v", err)
	}
	defer db.Close()

	store := NewUserStore(db)

	tests := []struct {
		des             string
		mockQueryReturn func()
		user            *models.Users
		expectedErr     error
	}{
		{
			des: "Success case",
			mockQueryReturn: func() {
				mock.ExpectExec(`INSERT INTO User (UserName, UserAge, Phone_number, Email) VALUES (?, ?, ?, ?)`).
					WithArgs("Chandana", 20, "9019878882", "chandana@gmail.com").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user: &models.Users{
				UserName:    "Chandana",
				UserAge:     20,
				Phonenumber: "9019878882",
				Email:       "chandana@gmail.com"},
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		test.mockQueryReturn()
		err := store.AddUser(test.user)
		assert.NoError(t, err)
	}
}

func Test_Deleteuser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error in sql mock %v", err)
	}
	defer db.Close()

	store := NewUserStore(db)
	mock.ExpectExec("DELETE FROM User WHERE UserName = ?").
		WithArgs("Chandana").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = store.Deleteuser("Chandana")
	assert.NoError(t, err)
}

func Test_UpdateEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error in sql mock %v", err)
	}
	defer db.Close()
	store := NewUserStore(db)

	mock.ExpectExec("UPDATE User SET Email = ? WHERE UserName = ?").
		WithArgs("Chandana@gmail.com", "Chandana").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = store.UpdateUserEmail("Chandana", "Chandana@gmail.com")
	assert.NoError(t, err)
}
