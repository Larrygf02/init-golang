package main

import "fmt"

type User struct {
	edad int
	nombre string
	apellido string
}

func main() {
	//var uriel User
	//fmt.Println(User{nombre: "Uriel", apellido: "Hernandez"})
	uriel := new(User)
	uriel.nombre = "Uriel"
	fmt.Println(uriel.nombre)
}

