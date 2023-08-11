package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tuannguyenandpadcojp/utils/valctx"
)

func main() {
	// parent context will cancel after 3 seconds
	parentCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	childCtx := context.WithValue(parentCtx, "name", "child")
	valCtx := valctx.ValueOnly(parentCtx)

	childCtxSignal := make(chan int)
	valCtxSignal := make(chan int)

	// Run the new goroutine with childCtx
	// We start counting from 1 to 10, but the parentCtx cancel after 3s
	// so this goroutine only count to 3 and will stop
	go func(ctx context.Context) {
		for i := 0; i < 10; i++ {
			if ctx.Err() != nil {
				break
			}
			fmt.Printf("childCtx Count :%d\n", i+1)
			time.Sleep(1 * time.Second)
		}

		fmt.Println("childCtx: Finished")
		childCtxSignal <- 1
	}(childCtx)

	// Run the new goroutine with valCtx
	// We start counting from 1 to 10, and the parentCtx cancel after 3s
	// but this goroutine still count to 10 because it does not effect by parentCtx
	go func(ctx context.Context) {
		for i := 0; i < 10; i++ {
			if ctx.Err() != nil {
				break
			}
			fmt.Printf("valCtx Count :%d\n", i+1)
			time.Sleep(1 * time.Second)
		}

		fmt.Println("valCtx: Finished")
		valCtxSignal <- 1
	}(valCtx)

	// wait for 2 child goroutines exit
	_, _ = <-childCtxSignal, <-valCtxSignal

	fmt.Println("Stop")
}
