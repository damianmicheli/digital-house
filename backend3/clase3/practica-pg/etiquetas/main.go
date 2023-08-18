package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	PrimerNombre string `json:"primer_nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"-"` // no exporta este campo
	Direccion string `json:"direccion,omitempty"` // ignora campo si esta vacio
}

func main() {
	p := Persona{"Juan", "Perez", "4343434", "Calle Falsa 123"}
	miJSON, err := json.Marshal(p)

	fmt.Println(string(miJSON))	
	if err != nil {fmt.Println(err)}

	p2 := Persona{
		PrimerNombre: "Juan",
		Apellido: "Carlos",
		Telefono: "12312334",
	}

	miJSON2, err2 := json.Marshal(p2)

	fmt.Println(string(miJSON2))	
	if err2 != nil {fmt.Println(err2)}
}