package main

import (
	"fmt"
	"sync"
)

func main() {
	RunNThreads(100)
}

func RunNThreads(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Thread %v\n", i)
		}(i)
	}
	wg.Wait()
}
