package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopping:", ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("working")
		}
	}

}

func main() {

	fmt.Println("worker start")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, &wg)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

	fmt.Println("worker end")

}
