package main

import (
	"github.com/assignment/handler"
	"github.com/assignment/service"
	"github.com/assignment/store"
	"gofr.dev/pkg/gofr"
)

func main() {

	app := gofr.New()
	
	store := store.New()
	ser := service.New(store)
	hand := handler.New(ser)

	app.GET("/users/{name}", hand.GetUserByName)
	app.POST("/users", hand.AddUser)
	app.DELETE("/users/{name}", hand.DeleteUser)
	app.PUT("/users/{name}", hand.UpdateEmail)

	app.Run()
}
