package main

import (
	"fmt"
)

func main() {
	defer function()
	integerNumb(1, 2, 3, 4, 5, 6, 7)
	var num = 13
	result, err := Fibonacci(num)
	if err != nil {
		fmt.Printf("This is the error: %s", err)
	} else {
		fmt.Printf("The Fibonacci value of %d is %d\n", num, result)
	}
	defer fmt.Println("======>", fibonacciEasy(15))

	i := 4
	withoutPointer(i)
	defer fmt.Println("withoutPointer", i)
	withPointer(&i)
	fmt.Println("withPointer", i)

	//For Advanced Function
	defer fmt.Printf("The type of this function is %T\n", returnParams)
	fmt.Printf("The type of this function for returnParams(4) %T\n", returnParams(4))
	fmt.Printf("Function as parameter %d\n", returnFunction(4, returnParams)) // refer Go Rough Doc

}

func function() {
	fmt.Printf("checking %s for the %d time. Now Time is %f\n", "printf", 1, 10.04)
}
func integerNumb(integer ...int) {
	for index, value := range integer {
		fmt.Printf("Index is %d: value is %d\n", index, value)
	}
}

func Fibonacci(fibNumb int) (int, error) {
	if fibNumb == 0 {
		return 0, nil
	} else if fibNumb == 1 {
		return 1, nil
	} else if fibNumb > 1 {
		result1, err1 := Fibonacci(fibNumb - 1)
		result2, err2 := Fibonacci(fibNumb - 2)
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("Error Calculating Fibonacci")
		} else {
			return result1 + result2, nil
		}
	} else {
		return 0, fmt.Errorf("Number should be non-negative number\n")
	}
}

func fibonacciEasy(fib int) int {
	if fib == 0 {
		return 0
	} else if fib == 1 {
		return 1
	} else {
		return fibonacciEasy(fib-1) + fibonacciEasy(fib-2)
	}
}

func withoutPointer(i int) {
	i = i + 1
}

func withPointer(i *int) {
	*i = *i + 1
}

// Advanced Function
func returnParams(i int) int {
	return i
}

// function as parameter
func returnFunction(i int, f func(int) int) int {
	return f(i)
}

// package main

// import (
// 	"fmt"
// )

// type MyStruct struct {
// 	Field1 int
// 	Field2 string
// }

// func (s MyStruct) String() string {
// 	return fmt.Sprintf("Field1: %d, Field2: %s", s.Field1, s.Field2)
// }

// func main() {
// 	myVar := MyStruct{Field1: 42, Field2: "Hello"}
// 	fmt.Printf("MyVar: %v\n", myVar)
// }
