package saludo

import "fmt"

//Saludo devuelve un saludo a la persona nombrada
func Hola(name string) string {
	// devuelve un saludo que incluye el nombre en el mensaje
	mensaje := fmt.Sprintf("Hola, %v. Bienvenido!", name)
	return mensaje
}