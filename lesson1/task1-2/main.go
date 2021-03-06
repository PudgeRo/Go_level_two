//go:build windows || linux
// +build windows linux

// If launch command "file main.exe": main.exe: PE32+ executable
// (console) x86-64 (stripped to external PDB), for MS Windows

package main

import (
	"fmt"
)

var (
	DivideByZero error = fmt.Errorf("cannot divide by zero")
)

func main() {
	fmt.Println(lastValue([]string{"one", "two", "three"}))

	result, err := Divide(5, 1)
	if err != nil {
		fmt.Println(fmt.Errorf("found error: %w", err))
		return
	}
	fmt.Println(result)
}

func lastValue(slice []string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("found error: %w", err))
		}
	}()

	return fmt.Sprintf("Last value is %v", slice[len(slice)])
}

// Divide returns a result of division number a by number b and error
func Divide(a, b float64) (float64, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("found error: %w", err))
		}
	}()

	if b == 0 {
		return 0, DivideByZero
	}
	return a / b, nil
}
