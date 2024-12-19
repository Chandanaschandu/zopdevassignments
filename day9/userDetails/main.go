package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"regexp"
)

// Users ,users is the struct
type Users struct {
	UserName     string `json:"user_name"`
	UserAge      int    `json:"user_age"`
	Phone_number string `json:"phone_Number"`
	Email        string `json:"email"`
}

// UsersList
type UsersList struct {
	db *sql.DB
}

func NewDetails(db *sql.DB) *UsersList {
	return &UsersList{db: db}
}

func (u *Users) Validate() error {

	emailRe := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	email := regexp.MustCompile(emailRe)
	isEmailValid := email.MatchString(u.Email)
	if !isEmailValid {
		return errors.New("Invalid email id")
	}

	phoneRe := `^\+?[0-9]{10,15}$`
	phone := regexp.MustCompile(phoneRe)
	isPhoneValid := phone.MatchString(u.Phone_number)
	if !isPhoneValid {
		return errors.New("Invalid phone number")
	}

	return nil

}

// GetUsers
func (userStore *UsersList) GetUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := userStore.db.Query("SELECT UserName,UserAge,Phone_number,Email FROM User")
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}
	var users = make([]Users, 0)
	for rows.Next() {
		var user Users
		_ = rows.Scan(&user.UserName, &user.UserAge, &user.Phone_number, &user.Email)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, "Error while iterating over the rows", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "error in marshalling userslist", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error in getting the user", http.StatusInternalServerError)
	}

}

// GetUsersByName
func (userStore *UsersList) GetUsersByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	var users Users
	query := "SELECT Username,UserAge,Phone_number,Email from User WHERE UserName=?"

	log.Printf("Executing query: %s", query)

	err := userStore.db.QueryRow(query, name).Scan(&users.UserName, &users.UserAge, &users.Phone_number, &users.Email)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error in querrying:"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error marshalling book: %v", err)
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Json data is not correct", http.StatusBadRequest)
	}

}

// AddUsers
func (userStore *UsersList) AddUsers(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	var NewUser Users
	err := json.Unmarshal(body, &NewUser)

	if err != nil {
		http.Error(w, "error in unmarshalling", http.StatusBadRequest)
	}

	if err := NewUser.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var exists bool
	_ = userStore.db.QueryRow("SELECT EXISTS(SELECT 1 from User where UserName=? AND UserAge=? AND Phone_number=?  And Email=?)", NewUser.UserName, NewUser.UserAge, NewUser.Phone_number, NewUser.Email).Scan(&exists)
	if exists {
		http.Error(w, "User already exits", http.StatusBadRequest)
		return
	}

	_, err = userStore.db.Exec("INSERT into User(UserName,UserAge,Phone_number,Email) VALUES (?,?,?,?)", NewUser.UserName, NewUser.UserAge, NewUser.Phone_number, NewUser.Email)
	if err != nil {
		http.Error(w, "Error in inserting new users to database", http.StatusBadRequest)
	}

	jsonData, _ := json.Marshal(NewUser)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Json data is incorrect", http.StatusBadRequest)
	}

}

// DeleteUsers
func (userStore *UsersList) DeleteUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Username := vars["name"]

	_, err := userStore.db.Exec("DELETE from User where UserName=?", Username)

	if err != nil {
		http.Error(w, "User name not found in database", http.StatusNotFound)
	}
	_, err = w.Write([]byte(" user Deleted sucessfully"))
	if err != nil {
		http.Error(w, "User not deleted", http.StatusBadRequest)
	}
}

// UpdateUsers
func (userStore *UsersList) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var users Users
	body, _ := io.ReadAll(r.Body)

	var updateUser Users
	err := json.Unmarshal(body, &updateUser)

	if err != nil {
		http.Error(w, "Error in unmarshalling", http.StatusInternalServerError)
	}

	if updateUser.Email != "" {
		users.Email = updateUser.Email
	}

	if err := users.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = userStore.db.Exec("UPDATE User SET Email=? WHERE Username=?", users.Email, name)
	if err != nil {
		http.Error(w, "Users doesnot exists", http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error in marshalling the body", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "json data is incorrect", http.StatusInternalServerError)
	}

}

// main function
func main() {

	db, err := sql.Open("mysql", "root:password@tcp/org_db")
	if err != nil {
		fmt.Println("Error", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database:%v", err)
	}

	fmt.Println("Successfully connected to the MySQL database!")

	userStore := NewDetails(db)
	//route ,creating route using mux
	route := mux.NewRouter()

	route.HandleFunc("/user", userStore.GetUsers).Methods("GET", "POST")
	route.HandleFunc("/user/{name}", userStore.GetUsersByName).Methods("GET")
	route.HandleFunc("/user", userStore.AddUsers).Methods("POST")
	route.HandleFunc("/user/{name}", userStore.UpdateUsers).Methods("PUT")
	route.HandleFunc("/user/{name}", userStore.DeleteUsers).Methods("DELETE")

	fmt.Println("Server started on http://localhost:8080")
	err = http.ListenAndServe(":8080", route)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
