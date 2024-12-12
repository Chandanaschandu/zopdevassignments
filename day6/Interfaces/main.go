package main

import "fmt"

type Calculator interface {
	Add()
	Sub()
	Mul()
	Div()
}
type Mystruct struct {
	num1 int
	num2 int
}

func (t Mystruct) Add() {
	fmt.Println("Addition of two numbers", t.num1+t.num2)
}
func (t Mystruct) Sub() {
	fmt.Println("Subtraction of two numbers", t.num1-t.num2)
}
func (t Mystruct) Mul() {
	fmt.Println("Multiplication of two numbers", t.num1*t.num2)
}
func (t Mystruct) Div() {
	fmt.Println("Division of Two numbers", t.num1/t.num2)
}

func main() {
	var i Calculator = Mystruct{2, 3}
	i.Add()
	i.Sub()
	i.Mul()
	i.Div()
}
