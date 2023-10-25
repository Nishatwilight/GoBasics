package main

import "fmt"

func main() {
	a := 3
	b := 6
	c := 12.4
	d := 5.6

	fmt.Println("Before swap")
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)

	swap(&a, &b)

	fmt.Println("After swap")
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)

	areaOfRectangle := calculateArea(&c, &d)
	fmt.Println("Area Of Rectangle: ", areaOfRectangle)

	numbers := []float64{5.2, 3.8, 7.4, 2.6}
	averageNumber := calculateAverage(&numbers)
	fmt.Println("Average of Numbers: ", averageNumber)

}

//Task 1: Swap Function
func swap(a *int, b *int) {
	// Your code to swap the values of a and b using pointers.
	*a, *b = *b, *a
}

//Task 2: Calculate the Area of a Rectangle
func calculateArea(length *float64, width *float64) float64 {
	answer := *length * *width
	return answer
}

//Task 3: Calculate the Average of Numbers
func calculateAverage(numbers *[]float64) float64 {
	lenofnum := len(*numbers)
	fmt.Println("lenofnum", lenofnum)
	var sum float64 = 0
	for _, singlenumber := range *numbers {
		sum += singlenumber
		fmt.Println("singlenumber", singlenumber)
	}
	fmt.Println("sum", sum)

	avg := sum / float64(lenofnum)
	return avg

}

//Task 4: Calculate the Maximum and Minimum

// func findMinMax(numbers *[]float64) (max *float64, min *float64) {
// }
