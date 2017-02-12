package main

import "fmt"

type Commodity string

const (
	GAS     Commodity = "gas"
	FOOD    Commodity = "food"
	CLOTHES Commodity = "clothes"
)

type actor struct {
	id          int
	product     Commodity
	commodities map[Commodity]int
}

type point struct {
	x int
	y Commodity
}

func main() {
	var ac = actor{id: 1, product: FOOD, commodities: map[Commodity]int{GAS: 1, FOOD: 2, CLOTHES: 0}}
	ac.id = 3
	fmt.Printf("\nHello, world: ", ac)

	var x = ac.id
	fmt.Printf("\nHello, world: ", x)
}
