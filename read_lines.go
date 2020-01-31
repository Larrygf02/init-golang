package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	archivo, error := os.Open("./hola.txt")
	defer archivo.Close()
	if error != nil {
		fmt.Println("hubo un error")
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}
	archivo.Close()
}