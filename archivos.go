package main

import (
	"fmt"
	"io/ioutil"
)
func main() {
	filedata, err := ioutil.ReadFile("./hola.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}
	fmt.Print(string(filedata))
}