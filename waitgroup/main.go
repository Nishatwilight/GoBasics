package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Lets learn Wait Group")
	var evilNinjas = []string{"ninja1", "ninja2", "ninja3"}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(evilNinjas))

	for _, ninjas := range evilNinjas {
		go attack(ninjas, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("Mission Completed")
}
func attack(evilninja string, waitGroup *sync.WaitGroup) {
	fmt.Println("Attacked eveil ninja is :", evilninja)
	waitGroup.Done()
}

//----------------------common Mistake In WeightGroup-----------------------------
// var weightGrp sync.WaitGroup
// weightGrp.Add(1)
// weightGrp.Wait()
// weightGrp.Done()
//----------Wrong Becoz the the order have to be Add() , Done(), Wait()-----------

//----------------------common Mistake In WeightGroup-----------------------------
// weightGrp.Add(1)
// weightGrp.Done()
// weightGrp.Done()
// weightGrp.Wait()
//--------Wrong Becoz the the order have to becoz 2 Done() are not allowed--------
