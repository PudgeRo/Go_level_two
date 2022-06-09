package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct{}

func (f myStruct) MarshalJSON() ([]byte, error) {
	return []byte(`{"a": 0}`), nil
}

func main() {
	j, _ := json.Marshal(myStruct{})
	fmt.Println(string(j))
}

//	Task1
//Для самостоятельного изучения. Исправить исходный код программы таким
//образом, чтобы команда vet не ругалась.
//
//Initial code
//
//package main
//import (
//"encoding/json"
//"fmt"
//)
//type MyStruct struct {}
//func (f MyStruct) MarshalJSON() (string, error) {
//	return `{"a": 0}`, nil
//}
//func main() {
//	j, _ := json.Marshal(MyStruct{})
//	fmt.Println(string(j))
//}

//	Task2
//Для самостоятельного изучения. Дополните исходный код примера из раздела golint
//документацией таким образом, чтобы линтер не ругался на её отсутствие у экспортируемых
//функций.
//If call golint forMyEducation.go then there will be warning
//forMyEducation.go:8:6: exported type MyStruct should have comment or be unexported
//forMyEducation.go:10:1: exported method MyStruct.MarshalJSON should have comment or be unexported
//
//Initial code
//
//package main
//
//import (
//"encoding/json"
//"fmt"
//)
//
//type MyStruct struct{}
//
//func (f MyStruct) MarshalJSON() ([]byte, error) {
//	return []byte(`{"a": 0}`), nil
//}
//
//func main() {
//	j, _ := json.Marshal(MyStruct{})
//	fmt.Println(string(j))
//}
