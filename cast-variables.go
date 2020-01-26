package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
	// castear entero a string
	edad_str := strconv.Itoa(edad)
	fmt.Println("Mi edad es "+edad_str)
}
