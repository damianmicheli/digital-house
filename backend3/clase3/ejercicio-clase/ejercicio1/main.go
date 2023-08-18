package main

import "fmt"

// Employee
// Una empresa necesita realizar una buena gestión de sus empleados, para esto
// realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos
// empleados. Los objetivos son:
// l) estructura Personcon los campos ID, Name, DateOfBirth
// 2) Crear una estructura Employeecon los campos: ID, y una composición
// con la estructura Person.
// 3) Realizar un método a la Employee que se llame lo
// que hará es realizar Ia impresión de los campos de un empleado.
// 4) Instanciar en la función main() tanto una Person como un Employee cargando sus
// resgx:tivos campos y, último, el método
// Si logramos realizar este pequeño programa, pg@imps ayudar a Ia empresa a solucionar Ia
// gestión de los empleados.

func main () {

	person1 := Person {
		ID: "1234",
		Name: "Carlos",
		DateOfBirth: "05/08/1980",
	}

	fmt.Println(person1.Name)

	employee1 :=  Employee {
		ID: "ABCD",
		Position: "Developer",
		Person: person1,
	}

	employee1.PrintEmployee()
}

type Person struct {
	ID string
	Name string
	DateOfBirth string
}

// paquete fecha
// go get google.golang.org/genproto/googleapis/type/date


//Embeber (Composicion)
type Employee1 struct {
	ID string
	Position string
	Person
}


// Definir un tipo

// Employee is a struct that represents an employee info.
type Employee struct {
	ID string
	Position string
	Person Person
}

func (e Employee) PrintEmployee(){
	fmt.Printf("Name: %s, Date of Birth: %s, Position: %s", e.Person.Name, e.Person.DateOfBirth, e.Position)
}
