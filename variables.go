package main

import "fmt"

func main() {
	// Las variables declaradas deben ser usadas sino 
	// go no compilará
	var edad int
	edad = 24
	// el operador := obvia el tipo de variable
	edad_anterior := 23
	edad_anterior = edad_anterior + edad
	name := "Raul"
	fmt.Println(name)
	fmt.Println(edad_anterior)
}