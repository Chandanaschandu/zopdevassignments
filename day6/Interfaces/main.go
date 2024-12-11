package main

import "fmt"

type I interface {
	Add()
	Sub()
	Mul()
	Div()
}
type T struct {
	a int
	b int
}

func (t T) Add() {
	fmt.Println("Addition of two numbers", t.a+t.b)
}
func (t T) Sub() {
	fmt.Println("Subtraction of two numbers", t.a-t.b)
}
func (t T) Mul() {
	fmt.Println("Multiplication of two numbers", t.a*t.b)
}
func (t T) Div() {
	fmt.Println("Division of Two numbers", t.a/t.b)
}

func main() {
	var i I = T{2, 3}
	i.Add()
	i.Sub()
	i.Mul()
	i.Div()
}
