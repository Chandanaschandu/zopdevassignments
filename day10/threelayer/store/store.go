package store

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Chandanaschandu/threelayer/models"
)

type store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) store {
	return store{db: db}
}

func (s *store) GetUserName(userName string) (*models.Users, error) {
	query := "SELECT UserName, UserAge, Phone_number, Email FROM User WHERE UserName = ?"
	row := s.db.QueryRow(query, userName)

	var user models.Users
	err := row.Scan(&user.UserName, &user.UserAge, &user.Phonenumber, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	return &user, nil
}

func (u store) AddUser(user *models.Users) error {
	_, err := u.db.Exec("INSERT INTO User (UserName, UserAge, Phone_number, Email) VALUES (?, ?, ?, ?)",
		user.UserName, user.UserAge, user.Phonenumber, user.Email)
	return err
}

func (u store) Deleteuser(name string) error {
	_, err := u.db.Exec("DELETE FROM User WHERE UserName = ?", name)
	return err
}

func (u store) UpdateUserEmail(name, email string) error {
	_, err := u.db.Exec("UPDATE User SET Email = ? WHERE UserName = ?", email, name)
	return err
}
