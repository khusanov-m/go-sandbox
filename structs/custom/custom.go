package custom

import (
	"fmt"
)

type str string 
// can be used to make aliases for types
// can be used to add method to normal types

// func (string) log() {} // this is not allowed on local data types

func (text str) log() {
	fmt.Println(text)
}

func Run() {
	var name str = "John"
  name.log()
}
