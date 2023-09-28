package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lets Learn GoRoutine")
	startTime := time.Now()
	fmt.Println("startTime", startTime)
	defer func() {
		fmt.Println("endTime", time.Since(startTime))
	}()
	evilNinjas := []string{"ninja1", "ninja2", "ninja3", "ninja4"}
	for _, targetNinja := range evilNinjas {
		go attack(targetNinja)
	}
	time.Sleep(time.Second * 2)

}
func attack(target string) {
	fmt.Printf("lets attack this %s ninja now\n", target)
	time.Sleep(time.Second)
}
