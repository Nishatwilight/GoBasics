package main

import (
	"fmt"
	"math/rand"
)

func main() {
	scorech := make(chan int)
	go channelIteration(scorech)
	for {
		message, open := <-scorech
		if !open {
			break
		}
		fmt.Printf("The score is : %d\n", message)
	}
	// for i := 1; i <= numberRounds; i++ {
	// 	fmt.Printf("The score is : %d\n", <-scorech)
	// }
}

func channelIteration(scorech chan int) {
	numberRounds := 3
	for i := 1; i <= numberRounds; i++ {
		// rand.Seed(time.Now().UnixNano())
		score := rand.Intn(10)
		scorech <- score // Send the integer score into the channel
	}
	close(scorech)
}
