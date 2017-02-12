//Executable commands must always use package main.
//There is no requirement that package names be unique across all packages linked into a single binary, 
// only that the import paths (their full file names) be unique.
package main

import (
	"testing"
	"fmt"
)

type (
	vehicle interface {
		speed() string
	}

	motorized struct {
		horsepower int
	}

	car struct {
		motorized
		brand string
		wheels int
		carspeed int
	}

	boat struct {
		motorized
		brand string
		boatspeed int
	}

	hybrid struct {
		car
		boat
	}

	//if language: string field were present in both, it would compile
	//but throw exception when you want to access the field me.language
	//should their types differ, this wouldn't compile

)

func (c car) speed() int {
	return c.carspeed
}

func (b boat) speed() int {
	return b.boatspeed
}

func TestDiamondStructures(t *testing.T) {

	h := hybrid{}
	fmt.Println(h) // {{{0}  0 0} {{0}  0}}

	h.boat.brand = "BMW"
	h.car.brand = "Mercedes"
	fmt.Println(h) //{{{0} Mercedes 0 0} {{0} BMW 0}}

	h.boatspeed = 10
	h.boat.boatspeed = 20
	fmt.Println(h.boatspeed) //20

	//Diamond problem - only runtime
	//fmt.Println(h.brand) //ambiguous selector h.brand
	//fmt.Println(h.speed()) //ambiguous selector h.speed
}
