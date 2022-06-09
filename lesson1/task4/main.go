package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}

// Initial function. It was necessary that there was no panic.
//func main() {
//	defer func() {
//		if v := recover(); v != nil {
//			fmt.Println("recovered", v)
//		}
//	}()
//	go func() {
//		panic("A-A-A!!!")
//	}()
//	time.Sleep(time.Second)
//}
