package main

import "fmt"

type person struct {
	name        string
	age         int
	phoneNumber string
	add         address
	//add *address
}

type address struct {
	city    string
	state   string
	pincode int
}

func main() {

	var personDetails person

	//var add address

	//personDetails.add = &add
	fmt.Println("Enter the person name")

	fmt.Scanln(&personDetails.name)
	fmt.Println("Enter the person age")

	fmt.Scanln(&personDetails.age)
	fmt.Println("Enter the person phone number")
	fmt.Scanln(&personDetails.phoneNumber)
	fmt.Println("Enter the person city")

	fmt.Scanln(&personDetails.add.city)
	fmt.Println("Enter the person state")
	fmt.Scanln(&personDetails.add.state)
	fmt.Println("Enter the person pincode")
	fmt.Scanln(&personDetails.add.pincode)

	fmt.Println(personDetails.name)
	fmt.Println(personDetails.age)
	fmt.Println(personDetails.phoneNumber)
	fmt.Println(personDetails.add.city)
	fmt.Println(personDetails.add.state)
	fmt.Println(personDetails.add.pincode)
	//fmt.Println(add.city)
}
