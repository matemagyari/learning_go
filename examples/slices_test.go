//Executable commands must always use package main.
//There is no requirement that package names be unique across all packages linked into a single binary, 
// only that the import paths (their full file names) be unique.
package main

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {

	//initializing flexible-length slice
	nums := []int{1, 2}

	for _, n := range nums {
		fmt.Println(n)
	}

	//would be runtime error
	//x1 := [2]int{1, 2, 3}
	//fmt.Println(x1)
}

func TestSlicesPassedByReference(t *testing.T) {

	ints := []int{1, 2}

	//leaves the original as it is
	ints2 := append(ints, 4)

	fmt.Println("original: ", ints)
	fmt.Println("updated: ", ints2)

	//changes the original
	changeValue := func(is []int) []int {
		is[0] = 999
		return is
	}

	updated := changeValue(ints)

	fmt.Println("original: ", ints)
	fmt.Println("updated: ", updated)
}

func TestArraysPassedByCopy(t *testing.T) {

	ints := [2]int{1, 2}

	//no append on an array - makes sense, it's fixed-length
	//ints2 := append(ints, 4)

	//does not change the original
	changeValue := func(is [2]int) [2]int {
		is[0] = 999
		return is
	}

	updated := changeValue(ints)

	fmt.Println("original: ", ints)
	fmt.Println("updated: ", updated)
}

func TestInitSlices(t *testing.T) {

	type person struct {
		name string
		age  int8
	}

	ps := []person{person{"joe", 21}, person{"jane", 18}}

	fmt.Println("original: ", ps)
}
