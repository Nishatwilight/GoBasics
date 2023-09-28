package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var treasureFound bool
var wg sync.WaitGroup
var count int = 100
var once sync.Once

func main() {
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			if FindTreasureNinjas() {
				once.Do(missionCompleted)
			}

			wg.Done()
		}()
	}
	wg.Wait()
	checkingCompleted()

}
func checkingCompleted() {
	if treasureFound {
		fmt.Println("Mission is Completed")
	} else {
		fmt.Println("Mission is Not Completed")
	}
}
func missionCompleted() {
	treasureFound = true
	fmt.Println("Mission Checking")

}
func FindTreasureNinjas() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
