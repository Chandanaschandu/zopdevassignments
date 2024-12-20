package store

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Chandanaschandu/threelayer/models"
)

type Store struct {
	db *sql.DB
}

func UserStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (u *Store) GetUserName(name string) (*models.Users, error) {
	var user models.Users

	err := u.db.QueryRow("SELECT UserName, UserAge, Phone_number, Email FROM User WHERE UserName = ?", name).
		Scan(&user.UserName, &user.UserAge, &user.Phonenumber, &user.Email)

	if err != nil {
		log.Fatalf("Error in getting the data")
	}

	return &user, err
}

func (u *Store) AddUser(user *models.Users) error {
	_, err := u.db.Exec("INSERT INTO User (UserName, UserAge, Phone_number, Email) VALUES (?, ?, ?, ?)",
		user.UserName, user.UserAge, user.Phonenumber, user.Email)
	return err
}

func (u *Store) Deleteuser(name string) error {
	_, err := u.db.Exec("DELETE FROM User WHERE UserName = ?", name)
	return err
}

func (u *Store) UpdateUserEmail(name, email string) error {
	_, err := u.db.Exec("UPDATE User SET Email = ? WHERE UserName = ?", email, name)
	return err
}
