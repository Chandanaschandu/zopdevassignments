package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Chandanaschandu/threelayer/handler"
	"github.com/Chandanaschandu/threelayer/service"
	"github.com/Chandanaschandu/threelayer/store"
)

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

	st := store.UserStore(db)
	ser := service.NewServices(st)
	userHandler := handler.NewUserHandler(ser)

	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", userHandler.GetUserByName).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/user", userHandler.AddUsers).Methods("POST")

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
