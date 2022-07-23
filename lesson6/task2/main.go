package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	fmt.Println("Before gosched")

	go showNumbers(10)

	//without Gosched goroutine showNumbers won't work
	runtime.Gosched()
	fmt.Println("After gosched")
}

func showNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
