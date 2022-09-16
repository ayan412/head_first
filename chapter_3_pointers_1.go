package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int // переменная	myIntPointer для хранения целочисленного указателя - тип указателя
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}
