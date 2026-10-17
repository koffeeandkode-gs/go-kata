package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {

	fmt.Println("before update", counter)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("after update", counter)
}
