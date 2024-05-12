// simple hello.go
package main

import (
	"fmt"
	"math"
)

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// with named return values
func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	var age int = 20 // with initial value
	age = 10         // reassignment
	var (
		name  string = "Hello"
		hours int    = int(math.Pow(10, 9))
	)
	coutnt := 10 // short-hand declaration
	// all types in go
	// bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128

	fmt.Println("Hello, World!", age, name, hours, coutnt)
	fmt.Printf("Hello, World! %d %s %d %d\n", age, name, hours, coutnt)

	type Person struct {
		name string
		age  int
	}

	var person Person = Person{name: "Hello", age: 20}
	fmt.Println(person)

	type myString string
	var str myString = "Hello"
	var str2 string = string(str) // we should convert myString to string before assign
	fmt.Println(str, str2)

	fmt.Println(calc(1, 2, "+"))

	fmt.Println(rectProps(1, 2))
	// the _ identifier is used to discard the perimeter.
	area, _ := rectProps2(1, 2)

	fmt.Println(area)

}

func calc(a int, b int, op string) int {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	}
	return a / b
}
