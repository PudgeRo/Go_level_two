package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)

	timer := time.After(time.Second)

	for {
		select {
		case <-timer:
			fmt.Println("time's up")
			return
		case sig := <-sigCh:
			fmt.Println("Stopped by signal: ", sig)
			return
		}
	}
}
