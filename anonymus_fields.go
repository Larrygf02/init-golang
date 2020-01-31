package main
import "fmt"

type Human struct {
	name string
	edad int
}

type Dummy struct {
	name string
}

type Tutor struct {
	human Human
	dummy Dummy
}

func main() {
	tutor := Tutor{Human{name:"Juan", edad: 12}, Dummy{"Mario"}}
	fmt.Println(tutor.human.name)
	fmt.Println(tutor.dummy.name)
	fmt.Println(tutor.human.edad)
}