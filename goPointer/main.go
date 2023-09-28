package main

import "fmt"

func main() {
	// Declare a variable 'x' of type int
	x := 10

	// Declare a pointer variable 'p' of type *int
	// and assign the address of 'x' to 'p'
	p := &x

	// Print the value of 'x' and the value pointed to by 'p'
	fmt.Println("Value of x:", x)
	fmt.Println("Value pointed to by p:", *p)

	// Modify the value of 'x' through the pointer 'p'
	*p = 20

	// Print the updated value of 'x' and the value pointed to by 'p'
	fmt.Println("Updated value of x:", x)
	fmt.Println("Value pointed to by p after modification:", *p)
}

//Golang is Case Sensitive, Check the below code:

// package main

// import "fmt"

// func main() {
// 	ab := "AB"
// 	AB := "AB"

// 	fmt.Println("The strings are not equal.", &ab)
// 	fmt.Println("The strings are not equal.", &AB)
// 	if ab != AB {
// 		fmt.Println("The strings are not equal.")

// 	}
// }
