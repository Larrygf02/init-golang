package main

import "fmt"

type User struct {
	edad int
	nombre string
	apellido string
}

func (usuario User) nombre_complet() string{
	return usuario.nombre + " " + usuario.apellido
}

// esta funcion no muta el objeto porque en la funcion
// hace una copia de la estructura
func (usuario User) set_name(n string) {
	usuario.nombre = n;
}

// este metodo si muta al estado ya que la estructura
// se pasa como punturo
func (usuario *User) set_namep(n string) {
	usuario.nombre = n;
}

func main() {
	//var uriel User
	//fmt.Println(User{nombre: "Uriel", apellido: "Hernandez"})
	uriel := new(User)
	uriel.nombre = "Uriel"
	uriel.apellido = "Gonzales"
	uriel.set_name("Marcos")
	fmt.Println(uriel.nombre_complet())
	fmt.Println(uriel.nombre)
}

