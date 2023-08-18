package main

import (
	"fmt"
	"errors"
)

/* Ejercicio
Crear un programa que cumpla los siguientes puntos:
1. Tener una estructura llamada Product con los campos ID, Name, Price,
Description y Category.
2. Tener un slice global de Product llamado Products, instanciado con
valores.
3. Dos métodos asociados a la estructura Product: Save(), GetAll(). El
método Save() deberá tomar el slice de Products y añadir el producto
desde el cual se llama al método. El método GetAll() deberá imprimir todos
los productos guardados en el slice Products.
4. Una función getById() a la cual se le deberá pasar un INT como parámetro
y retorna el producto correspondiente al parámetro pasado.
5. Ejecutar al menos una vez cada método y función definidos desde main().*/

// **** Declaraciones globales *****

var (
	prod1 = Product {1234, "Pelota", 5000.0, "Una pelota de futbol", "Articulos deportivos"}
	prod2 = Product {4567, "Raqueta", 15000.0, "Una raqueta de tenis","Articulos deportivos"}
	prod3 = Product {8901, "Antiparras", 3000.0, "Antiparras para nadar", "Articulos deportivos"}
	
	Products = []Product{prod1, prod2, prod3}
)

// ***** fin declaraciones globales *****


func main () {

	var p Product // genero un Product vacio solo para poder llamar al getall
	fmt.Println("\nSlice Original:")
	p.GetAll()

	prod4 := Product {2345, "Palo de golf", 30000.0, "Un palo para jugar al golf", "Articulos deportivos"}

	prod4.Save()
	
	fmt.Println("\nSlice Despues del Save:")

	p.GetAll()

	fmt.Println("\nBuscar por ID existente")

	productoBuscado, error := getById(4567)

	if error == nil {
		fmt.Println(productoBuscado.toString())
	} else {
		fmt.Println(error)
	}
	
	fmt.Println("\nBuscar por ID inexistente")

	productoBuscado1, error1 := getById(44567)

	if error1 == nil {
		fmt.Println(productoBuscado1.toString())
	} else {
		fmt.Println(error1)
	}

}

type Product struct {
	ID int
	Name string
	Price float64
	Description string
	Category string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) toString() string {
	return fmt.Sprintf("ID: %d \n Nombre: %s \n Precio: %.2f \n Descripcion: %s \n Categoria: %s", 
				p.ID, p.Name, p.Price, p.Description, p.Category)
}

//Imprime todos los productos guardados en el slice Products.
func (p Product) GetAll() {
	for _, item := range Products {
		fmt.Println(item.toString())
	}

}

//Retorna el producto correspondiente al INT pasado como parámetro.
func getById(id int) (Product, error){

	var encontrado Product = Product{ID: -1}

	for _, product := range Products {
		if product.ID == id {
			encontrado = product
			break
		}
	}

	if encontrado.ID == -1 {
		return encontrado, errors.New("no existe producto con el ID indicado")
	}

	return encontrado, nil
}


