package main

import "fmt"

func main() {
	ab := "AB"
	AB := "AB"

	fmt.Println("The strings are not equal.", &ab)
	fmt.Println("The strings are not equal.", &AB)
	if ab != AB {
		fmt.Println("The strings are not equal.")

	}
}
