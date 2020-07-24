package main

import "fmt"

type salarized interface {
	getSalary() int
}

type Salary struct {
	basic     int
	allowance int
	insurance int
}

func (s *Salary) getSalary() int {}

type Employee struct {
	firstName string
	lastName  string
	salary    Salary
	fullTime  bool
}

func main() {
	ross := &Employee{
		firstName: "Vu",
		lastName:  "Tran",
		salary:    Salary{1100, 50, 50},
		fullTime:  true,
	}

	fmt.Println(ross)

}
