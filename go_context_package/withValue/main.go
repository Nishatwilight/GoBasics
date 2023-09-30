package main

import (
	"context"
	"fmt"
)

func myfunc(ctx context.Context) {
	fmt.Printf("myfunc: key's value is %s\n", ctx.Value("key"))
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "scaler")

	myfunc(ctx)

}
