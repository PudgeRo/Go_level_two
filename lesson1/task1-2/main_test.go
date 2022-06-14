package main

import "fmt"

func ExampleDivide() {
	result, err := Divide(10, 5)
	if err != nil {
		fmt.Println("error has occurred")
	}
	fmt.Println(result)
}
