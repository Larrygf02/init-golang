package main

import (
	"fmt"
)

func main() {
	fmt.Println("Comprueba de dos numeros si es mayor menor o igual")
	fmt.Println("==================================================")
	fmt.Println("Ingresa un numero: ")
	var numberOne, numberTwo int
	fmt.Scanf("%d\n", &numberOne)
	fmt.Println("Ingresa otro numero: ")
	fmt.Scanf("%d\n", &numberTwo)

	if numberOne == numberTwo {
		fmt.Printf("%d es igual a %d\n", numberOne, numberTwo)
	}else if numberOne > numberTwo {
		fmt.Printf("%d es mayor que %d\n", numberOne, numberTwo)
	}else {
		fmt.Printf("%d es menor que %d\n", numberOne, numberTwo)
	}
}