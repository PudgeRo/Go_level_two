package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	var mu sync.Mutex

	sum := 0
	workers := make(chan struct{}, runtime.NumCPU())
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-workers
			}()
			mu.Lock()
			sum++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
