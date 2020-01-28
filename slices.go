// Los slices nos permite redimensionar a diferencia de los arreglos
package main
import "fmt"

func main() {
	// slices := []int{1,2,3,4}
	// fmt.Println(len(slices))
	slice := make([]int, 3, 5)
	slice = append(slice, 2)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	sliceone := []int{1,2,3,4}
	copia := make([]int, len(slice))

	copy(copia, sliceone)

	fmt.Println(copia)
}