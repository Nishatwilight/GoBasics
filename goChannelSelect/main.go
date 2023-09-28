package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Channel Select Statement")
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Number 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- true
	}()
	time.Sleep(2 * time.Second)

	select {
	case chan1 := <-ch1:
		fmt.Printf("The integer value is: %d\n", chan1)
	case chan2 := <-ch2:
		fmt.Printf("The String value is: %s\n", chan2)
	case chan3 := <-ch3:
		fmt.Printf("The Boolean value is: %t\n", chan3)
	default:
		fmt.Println("The channel is not ready")
	}
	roughlyFair()
}

func roughlyFair() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)
	ninja3 := make(chan interface{})
	close(ninja3)
	ninja4 := make(chan interface{})
	close(ninja4)

	var ninjacount1, ninjacount2, ninjacount3, ninjacount4 int

	for i := 1; i <= 1000; i++ {
		select {
		case <-ninja1:
			ninjacount1++
		case <-ninja2:
			ninjacount2++
		case <-ninja3:
			ninjacount3++
		case <-ninja4:
			ninjacount4++
		default:
			fmt.Println("The counting process is going on")
		}

	}
	fmt.Printf("the count of ninjacount1 : %d\n the count of ninjacount2 : %d\n the count of ninjacount3 : %d\n the count of ninjacount4 : %d\n", ninjacount1, ninjacount2, ninjacount3, ninjacount4)
}

// func captionElect(ninja chan string, message string) {
// 	ninja <- message
// }
