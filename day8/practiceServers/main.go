package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Employee is the Main struct
type Employee struct {
	EmployeeId     int    `json:"employee_id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
}

// DBEmployee is the struct which has database sql connection
type DBEmployee struct {
	db *sql.DB
}

// NewStore create newemployee
func NewEmployee(db *sql.DB) *DBEmployee {
	return &DBEmployee{db: db}
}

func (employee *DBEmployee) GetEmployee(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request, book ID missing", http.StatusBadRequest)
		return
	}
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}
	var employee1 Employee
	query := "SELECT EmployeeId, EmployeeName, EmployeeSalary FROM Employee WHERE EmployeeId= ?"
	log.Printf("Executing this query", query)
	err = employee.db.QueryRow(query, id).Scan(&employee1.EmployeeId, &employee1.EmployeeName, &employee1.EmployeeSalary)
	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {

		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(employee)
	if err != nil {
		log.Printf("Error marshalling employee: %v", err)
		http.Error(w, "Error marshalling employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (employee *DBEmployee) AddEmployee(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var employee1 Employee
	_ = json.Unmarshal(body, &employee1)

	var exists bool
	_ = employee.db.QueryRow("SELECT EXISTS(SELECT 1 FROM Employee WHERE EmployeeName = ? AND EmployeeSalary = ?)", employee1.EmployeeName, employee1.EmployeeSalary).Scan(&exists)

	if exists {
		http.Error(w, "The employee alreadye exists", http.StatusBadRequest)
		return
	}

	// Insert the new book
	//result, _ := employee.db.Exec("INSERT INTO Employee (EmployeeName, EmployeeSalary) VALUES (?, ?)", employee1.EmployeeName, employee1.EmployeeSalary)
	result, err := employee.db.Exec("INSERT INTO Employee (EmployeeName, EmployeeSalary) VALUES (?, ?)", employee1.EmployeeName, employee1.EmployeeSalary)
	if err != nil {
		http.Error(w, "Error inserting employee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	employee1.EmployeeId = int(id)

	jsonData, _ := json.Marshal(employee1)
	w.Write(jsonData)
}
func (employee *DBEmployee) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, _ := strconv.Atoi(idStr)

	_, _ = employee.db.Exec("DELETE FROM Employee WHERE EmployeeId= ?", id)

	response := map[string]string{
		"message": "Employee deleted",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)

}
func main() {
	db, err := sql.Open("mysql", "root:password@tcp/org_db")

	if err != nil {
		fmt.Printf("Error in open mysql %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("The database connected sucessfully")

	employee := NewEmployee(db)
	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			path := r.URL.Path
			parts := strings.Split(path, "/")
			if len(parts) > 1 && parts[len(parts)-1] != "" {

				employee.GetEmployee(w, r)
			}
		} else if r.Method == http.MethodPost {
			employee.AddEmployee(w, r)
		} else if r.Method == http.MethodDelete {
			employee.DeleteEmployee(w, r)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("server successfully started")
}
