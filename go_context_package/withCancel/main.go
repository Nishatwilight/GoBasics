package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(2*time.Second, cancel)

	sayMyName(ctx, 6*time.Second, "scaler")
}

func sayMyName(ctx context.Context, d time.Duration, name string) {
	select {
	case <-time.After(d):
		fmt.Print("Your name is ", name)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Print(err)
	}
}
