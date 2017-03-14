package main

import (
	"fmt"
	"testing"
)

type point struct {
	x, y int
}

func changeValue(p point) {
	p.x = p.x * 2
}

func changeRef(p *point) {
	p.x = p.x * 2
}

func TestPointersForStructs(t *testing.T) {

	p := point{33, 4}
	changeValue(p)
	fmt.Println(p.x) //33
	changeRef(&p)
	fmt.Println(p.x) //66
}


func TestPointersForSlices(t *testing.T) {

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
