package main

import (
	"fmt"
)

type employee interface {
	calculateSalary() int
}

type Fulltime struct {
	id int
	days int
	basic int
}

func (ft *Fulltime) calculateSalary() int {
	return ft.basic * ft.days
}

type Contractor struct {
	id int
	days int
	basic int
}

func (ct *Contractor) calculateSalary() int {
	return ct.basic * ct.days
}

type Freelancer struct {
	id int
	hours int
	basic int
}

func (fl *Freelancer) calculateSalary() int {
	return fl.basic * fl.hours
}

func main() {
	var em employee

	em = &Fulltime{1, 28, 500}
	fmt.Println(em.calculateSalary())

	em = &Contractor{2, 28, 100}
	fmt.Println(em.calculateSalary())

	em = &Freelancer{3, 45, 10}
	fmt.Println(em.calculateSalary())

}