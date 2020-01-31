package main
import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre string
	numero_permiso int
}

func (admin Admin) Permisos() int{
	return admin.numero_permiso
}

func (admin Admin) Nombre() string{
	return admin.nombre
}
 
func auth(user User) string {
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de admin"
	}
	return "No eres admin"
}

func main() {
	admin := Admin{nombre: "Raul", numero_permiso: 5}
	fmt.Println(auth(admin))
}