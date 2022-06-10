package main

import (
	"fmt"
	"log"
	"os"
)

var (
	PanicEveryTenthFile = fmt.Errorf("panic every tenth file")
	ErrorCreateFile     = fmt.Errorf("cannot create file")
	ErrorCloseFile      = fmt.Errorf("cannot close file")
)

// Function creates files and panic every tenth files
func createWithPanic(name string, i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("%w", err))
			file, err := os.Create("./files/" + name)
			if err != nil {
				log.Fatal(ErrorCreateFile)
			}
			fmt.Printf("%v has been created\n", name)
			err = file.Close()
			if err != nil {
				log.Fatal(ErrorCloseFile)
			}
			fmt.Printf("%v has been closed\n", name)
		}
	}()
	if i%10 == 0 {
		panic(PanicEveryTenthFile)
	} else {
		file, err := os.Create("./files/" + name)
		if err != nil {
			log.Fatal(ErrorCreateFile)
		}
		fmt.Printf("%v has been created\n", name)
		err = file.Close()
		if err != nil {
			log.Fatal(ErrorCloseFile)
		}
		fmt.Printf("%v has been closed\n", name)
	}
}
func createFiles(quantity int) {
	for i := 1; i <= quantity; i++ {
		name := "file " + fmt.Sprintf("%v.txt", i)
		createWithPanic(name, i)
	}
}

func main() {
	createFiles(100)
}
