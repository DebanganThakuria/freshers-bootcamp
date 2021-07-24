package main

type empSalCalc interface {
	calculateSalary(a, b int) int
}

type emp struct{}

func (c emp) calculateSalary(a, b int) int {
	return a * b
}

type Fulltime struct {
	id    int
	days  int
	basic int
	emp
}

type Contractor struct {
	id    int
	days  int
	basic int
	emp
}

type Freelancer struct {
	id    int
	hours int
	basic int
	emp
}

//func main() {
//	var esc empSalCalc
//
//	esc = Fulltime{1, 28, 500, emp{}}
//	fmt.Println(esc.calculateSalary(esc.(Fulltime).days, esc.(Fulltime).basic))
//
//	esc = Contractor{2, 28, 100, emp{}}
//	fmt.Println(esc.calculateSalary(esc.(Contractor).days, esc.(Contractor).basic))
//
//	esc = Freelancer{3, 45, 10, emp{}}
//	fmt.Println(esc.calculateSalary(esc.(Freelancer).hours, esc.(Freelancer).basic))
//}
