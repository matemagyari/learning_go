package main

import (
	"fmt"
	"testing"
)

type dog struct {
	brand string
}

func TestVariables(t *testing.T) {

	const name = "Joe"
	fmt.Println(name) //Joe

	//name = "" //doesn't compile

	//initialized to default value
	var age int
	var d dog

	fmt.Println(age) //0
	fmt.Println(d) //{}

	var i1 int = 4
	var i2 = 4
	var i4, i5 int = 3, 4
	i3 := 4
	//The := syntax is shorthand for declaring and initializing a variable
	i6, i7 := 4, 5
	fmt.Println(i1, i2, i3, i4, i5, i6, i7)

	age = 2
	fmt.Println(age) //2

	height := 4
	fmt.Println(height) //4

	height = 5
	fmt.Println(height) //5
}
