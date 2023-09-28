package main

import (
	"fmt"
)

//	func main() {
//		venue := []string{"school", "college", "Work", "Home", "office"}
//		fmt.Println("venue", venue)
//		for _, venueList := range venue {
//			if venueList == "school" {
//				greeting("i am in school")
//			} else if venueList == "Work" {
//				greeting("i am in Work")
//			} else if venueList == "Home" {
//				greeting("i am in Home")
//			} else if venueList == "office" {
//				greeting("i am in office")
//			} else if venueList == "college" {
//				greeting("i am in college")
//			}
//		}
//	}
func main() {
	// venue := []string{""}
	venue := []string{"school", "college", "Work", "Home", "office"}
	fmt.Println("venue", venue)
	for _, venueList := range venue {
		switch venueList {
		case "Home":
			greeting("i am in Home")
		case "office":
			greeting("i am in office")
		case "Work":
			greeting("i am in Work")
		case "school":
			greeting("i am in school")
		case "college":
			greeting("i am in college")
		default:
			greeting("i am noWhere")

		}
	}
}
func greeting(greet string) {
	fmt.Println(greet)
}
