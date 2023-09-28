package main

import "fmt"

type student struct {
	name string
}
type domain struct {
	department string
}

type check interface {
	count()
}

func main() {
	fmt.Println("Type Assertion")

	std := []check{
		student{name: "Nisha"},
		domain{department: "Mechanical"},
	}
	for _, students := range std {
		students.count()
		ns, ok := students.(student)
		if ok {
			ns.feesPaid()
		}
	}
}

func (d domain) count() {
	fmt.Println("Printing Domain", d.department)
}
func (s student) count() {
	fmt.Println("Printing Student", s.name)
}
func (f student) feesPaid() {
	fmt.Println("fees paid")
}
