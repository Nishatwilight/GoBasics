package main

import "fmt"

func main() {
	type ninja struct {
		name   string
		wepons []string
		level  int
	}

	ninjaMembers := ninja{name: "nisha"}
	fmt.Println("name", ninjaMembers)
	ninjaMembers = ninja{name: "jeev", wepons: []string{"Ak-47"}, level: 1}
	fmt.Println("name", ninjaMembers)
	ninjaMembers.level++
	ninjaMembers.wepons = append(ninjaMembers.wepons, "sniper", "short gun")
	fmt.Println("name", ninjaMembers)

	type dojo struct {
		name      string
		ninjaName ninja
	}
	golangDojo := dojo{
		name:      "Golang Dojo",
		ninjaName: ninjaMembers,
	}
	fmt.Println("ninjaMembers.level", ninjaMembers.level)
	fmt.Println("golangDojo", golangDojo)

	golangDojo.ninjaName.level = 5
	fmt.Println("golangDojo After level change", golangDojo)

	type addressDojo struct {
		name        string
		ninjaMember *ninja
	}

	address := addressDojo{name: "Address GolangDojo", ninjaMember: &ninjaMembers}
	fmt.Println("address", address)
	fmt.Println("*address.ninjaMember", *address.ninjaMember)
	address.ninjaMember.level = 7
	fmt.Println("address.ninjaMember.level check", *address.ninjaMember)

	jonnyPointer := new(ninja)
	jonnyPointer.name = "nisha"
	fmt.Println("jonnyPointer", jonnyPointer)
	fmt.Println("*jonnyPointer", *jonnyPointer)

	//How to add Func into the struct?

	ninjass := ninjaIntern{name: "internOne", level: 3}
	ninjass.sayHello()
	ninjass.sayName()
	fmt.Println("ninja", ninjass)
	fmt.Println(" ninjass.name", ninjass.name)

}

type ninjaIntern struct {
	name  string
	level int
}

func (ninjaIntern) sayHello() {
	fmt.Println("Hello")
}

func (n *ninjaIntern) sayName() {
	fmt.Println("Name of the ninjaIntern", n.name)
	n.name = "internTwo"
	fmt.Println("Name of the ninjaIntern", n.name)
}
