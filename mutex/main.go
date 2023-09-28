package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	lock   sync.Mutex
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func main() {
	fmt.Println("Lets learn Mutex")
	basics()
	readWriteMutex()
	wg.Wait()
}

func basics() {
	iterationCount := 1000
	for i := 0; i < iterationCount; i++ {
		go iteration()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Total Count is: ", count)
}
func iteration() {
	lock.Lock()
	count++
	lock.Unlock()
}

func readWriteMutex() {
	wg.Add(3)
	go read()
	go read()
	go write()
}

func read() {
	defer wg.Done()
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Println("Locking the Read Mutex")
	time.Sleep(2 * time.Second)
	fmt.Println("UnLocking the Read Mutex")
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	defer rwLock.Unlock()
	fmt.Println("Locking the Write Mutex")
	time.Sleep(2 * time.Second)
	fmt.Println("UnLocking the Write Mutex")
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	data        int
// 	dataMutex   sync.Mutex   // Basic Mutex to protect data
// 	dataRWMutex sync.RWMutex // RWMutex to protect data
// )

// func main() {
// 	// Start multiple reader goroutines
// 	for i := 0; i < 5; i++ {
// 		go reader(i)
// 	}

// 	// Start a writer goroutine
// 	go writer()

// 	// Allow some time for goroutines to run
// 	time.Sleep(3 * time.Second)
// }

// func reader(id int) {
// 	dataRWMutex.RLock() // Read lock
// 	defer dataRWMutex.RUnlock()

// 	fmt.Printf("Reader %d: Reading data: %d\n", id, data)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Reader %d: Finished reading\n", id)
// }

// func writer() {
// 	dataRWMutex.Lock() // Exclusive lock for writing
// 	defer dataRWMutex.Unlock()

// 	fmt.Println("Writer: Writing data...")
// 	data = 42 // Simulating a write operation
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Writer: Finished writing")
// }
