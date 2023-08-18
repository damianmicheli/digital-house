package main

import "fmt"

/*
Una empresa de redes sociales requiere implementar una estructura usuarios con
funciones que vayan agregando información a la misma. Para optimizar y ahorrar memoria
requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del
programa y para las funciones. La estructura debe tener los campos: nombre, apellido,
edad, correo y contraseña. Y deben implementarse las funciones:
● cambiarNombre: permite cambiar el nombre y apellido.
● cambiarEdad: permite cambiar la edad.
● cambiarCorreo: permite cambiar el correo.
● cambiarContraseña: permite cambiar la contraseña.
*/


type usuario struct {
	nombre string
	apellido string
	edad int
	correo string
	contrasena string
}

func (u usuario) detalles(){
	fmt.Printf("\nNombre: %s\nApellido: %s\nEdad: %d\nCorreo: %s\nContraseña: %s\n", 
	u.nombre, u.apellido, u.edad, u.correo, u.contrasena)
}

func crearUsuario(nombre, apellido string, edad int, correo, password string) usuario{
	return usuario {nombre, apellido, edad, correo, password}
}

func (u *usuario) cambiarNombre(nuevoNombre string) {
	u.nombre = nuevoNombre
}

func (u *usuario) cambiarEdad(nuevaEdad int) {
	u.edad = nuevaEdad
}

func (u *usuario) cambiarCorreo(nuevoCorreo string) {
	u.correo = nuevoCorreo
}

func (u *usuario) cambiarContrasena(nuevaContrasena string) {
	u.contrasena = nuevaContrasena
}


func main() {

	usr1 := crearUsuario("Damian", "Micheli", 43, "damianmicheli@gmail.com", "incorrecta")

	usr1.detalles()
	
	usr1.cambiarNombre("Juan")
	usr1.cambiarEdad(50)
	usr1.cambiarCorreo("juanm@gmail.com")
	usr1.cambiarContrasena("password")
	
	
	usr1.detalles()
}