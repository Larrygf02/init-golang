package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
/* 	edad := 22
	fmt.Printf("Mi edad es %d\n", edad)
	precio := 3.2
	fmt.Printf("El pan cuesta %f\n", precio)
	var curso int
	fmt.Println("Coloca numero de cursos: ")
	fmt.Scanf("%d\n", &curso)
	fmt.Printf("Estudiaré %d cursos", curso) */
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre , err := reader.ReadString('\n')
	if (err != nil) {
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+ nombre)
	}
}