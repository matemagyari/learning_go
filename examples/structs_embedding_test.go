//Executable commands must always use package main.
//There is no requirement that package names be unique across all packages linked into a single binary, 
// only that the import paths (their full file names) be unique.
package main

import (
	"fmt"
	"testing"
)

type (
	person interface {
		greet() string
	}

	engineer struct {
		id   int
		name string
	}

	computerEngineer struct {
		engineer
		language string
	}

	electricalEngineer struct {
		engineer
		field string
	}
)

func (e engineer) greet() string {
	return fmt.Sprintf("engineer %s", e.name)
}

func (e computerEngineer) greet() string {
	return fmt.Sprintf("computerEngineer %s", e.name)
}

func TestEmbeddedStructures(t *testing.T) {
	fmt.Println(engineer{id: 0, name: "Joe"}.greet())//engineer Joe

	ce := computerEngineer{engineer: engineer{id: 0, name: "Joe"}, language: "Go"}
	ee := electricalEngineer{engineer: engineer{id: 0, name: "Joe"}, field: "Powerplants"}

	//fields are initialized to default value

	//function is overridden
	fmt.Println(ce.greet()) //computerEngineer Joe
	//function is "inherited" from the embedded struct
	fmt.Println(ee.greet()) //engineer Joe

	//structures are mutable
	ce.language = "Cobol"
	fmt.Println(ce.language) // Cobol
}
