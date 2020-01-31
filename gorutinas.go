package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	// la palabra reservada go hace que la funcion sea concurrente
	go mi_nombre_lento("Raul hola como estas")
	fmt.Println("Que aburridoo")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lento(name string) {
	letras := strings.Split(name,"")
	for _, letra := range(letras) {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}