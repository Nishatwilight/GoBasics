package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lets Learn Go Channels")
	startTime := time.Now()
	// fmt.Println("startTime", startTime)
	defer func() {
		fmt.Println("endTime", time.Since(startTime))
	}()

	evilNinjas := []string{"ninja1", "ninja2", "ninja3", "ninja4"}
	for _, targetNinja := range evilNinjas {
		smokeSignalch := make(chan bool)
		go attack(targetNinja, smokeSignalch)
		fmt.Println("smokeSignalch", <-smokeSignalch)
		time.Sleep(time.Second * 2)
	}

}
func attack(target string, smokeSignalch chan bool) {
	time.Sleep(time.Second)
	fmt.Printf("lets attack this %s  now\n", target)
	smokeSignalch <- true
}
