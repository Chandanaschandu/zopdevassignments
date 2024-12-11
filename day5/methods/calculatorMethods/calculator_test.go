package calculatorMethods

import (
	"testing"
)

func TestAdd(t *testing.T) {
	c := Calculator{a: 3, b: 2}
	result := c.add()
	expected := 5.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
func TestSub(t *testing.T) {
	c := Calculator{a: 5, b: 3}
	result := c.sub()
	expected := 2.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
func TestMul(t *testing.T) {
	c := Calculator{a: 4, b: 3}
	result := c.mul()
	expected := 12.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
func TestDiv(t *testing.T) {
	c := Calculator{a: 6, b: 2}
	result := c.div()
	expected := 3.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
func TestDivByZero(t *testing.T) {
	c := Calculator{a: 6, b: 0}
	result := c.div()
	expected := 0.0
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
func BenchmarkAdd(b *testing.B) {
	c := Calculator{a: 3, b: 2}
	for range b.N {
		c.add()
	}
}
func BenchmarkSub(b *testing.B) {
	c := Calculator{a: 3, b: 2}
	for range b.N {
		c.sub()
	}
}
func BenchmarkMul(b *testing.B) {
	c := Calculator{a: 3, b: 2}
	for range b.N {
		c.mul()
	}
}
func BenchmarkDiv(b *testing.B) {
	c := Calculator{a: 3, b: 2}
	for range b.N {
		c.div()
	}
}
