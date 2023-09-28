package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	ctx := context.Background()
	userID := 10
	value, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %d\n", value)
	defer fmt.Println("Took us: ", time.Since(start))

}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, UserID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)
	go func() {
		val, err := fetchThirdParty()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}
func fetchThirdParty() (int, error) {
	time.Sleep(time.Millisecond * 125)
	return 666, nil
}
