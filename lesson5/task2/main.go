package main

import (
	"fmt"
	"sync"
)

const n = 100

func main() {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		counter int
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
