package main

import "fmt"

func main() {
	fmt.Println("Buffered Channel")

	firstCh := make(chan string, 3)
	// go func() {
	firstCh <- "First Message"
	firstCh <- "Second Message"

	// }()
	fmt.Println("firstCh: ", <-firstCh)
	fmt.Println("firstCh: ", <-firstCh)

}
