package main

import "fmt"

type (
	employee interface {
		yearContribution() float32
		baseSalary() int
		toString() string
	}

	baseEmployee struct {
		name  string
		years int
	}

	officeWorker struct {
		baseEmployee
	}

	fieldWorker struct {
		baseEmployee
		physicalWork bool
	}
)

func (_ officeWorker) baseSalary() int {
	return 500
}

func (_ fieldWorker) baseSalary() int {
	return 1000
}

func (o officeWorker) yearContribution() float32 {
	return 30 * float32(o.years)
}

func (f fieldWorker) yearContribution() float32 {

	if f.physicalWork {
		return 40 * float32(f.years)
	} else {
		return 30 * float32(f.years)
	}
}

func salary(e employee) float32 {
	return float32(e.baseSalary()) + e.yearContribution()
}

func (e baseEmployee) toString() string {
	return fmt.Sprintf("My name is %s and I've been working here for %d years.", e.name, e.years)
}

func main() {

	ow := officeWorker{baseEmployee{name:"joe", years:20}}
	fw := fieldWorker{baseEmployee: baseEmployee{name:"jack", years:20}, physicalWork: true}

	for _, e := range []employee{ow, fw} {
		fmt.Println(e.toString())
		fmt.Println(salary(e))
	}
}