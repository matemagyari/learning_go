package main

import (
	"fmt"
	"testing"
)

type point struct {
	x, y int
}

func change(p point) {
	p.x = p.x * 2
}

func change2(p *point) {
	p.x = p.x * 2
}

func TestPointers(t *testing.T) {

	p := point{33, 4}
	change(p)
	fmt.Println(p.x) //33
	change2(&p)
	fmt.Println(p.x) //66
}
