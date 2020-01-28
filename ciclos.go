package main
import "fmt"

func main() {
/* 	for i:=0; i<10;i++ {
		fmt.Println("Hola ciclo")
	} */
	i:= 0
	for {
		if i == 2{
			i++
			continue // regresa al ciclo for
		}
		fmt.Println(i)
		i++

		if i>10{
			break
		}
	}
}