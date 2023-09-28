package main

import (
	"fmt"
	"time"
)

func main() {
	maxRetries := 5
	retryDelay := 2 * time.Second
	for retries := 0; retries < maxRetries; retries++ {
		if err := performOperation(); err != nil {
			fmt.Printf("Error: %v. Retrying...\n", err)
			time.Sleep(retryDelay)
			retryDelay *= 2 // Exponential backoff
			continue
		}
		fmt.Println("Operation successful.")
		break
	}
}
func performOperation() error {
	// Simulate an operation that may fail
	return nil // Replace with your actual operation
}
