package main

import (
	"errors"
	"fmt"
)

/*
Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total. La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
Pequeño: solo tiene el costo del producto.
Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total basándose en el costo del producto y los adicionales, en caso de que los tenga.
*/
const (
	smallSize = "SMALL"
	mediumSize = "MEDIUM"
	largeSize = "LARGE"
	shipping = 2500.0
	mediumStorage = 3
	largeStorage = 6
)

type product interface {
	getPrice() float64
}

type small struct {
	name string
	description string
	unitCost float64
}

func (s small) getPrice() float64 {
	return s.unitCost
}

type medium struct {
	name string
	description string
	unitCost float64
}

func (m medium) getPrice() float64 {
	return m.unitCost * (1 + mediumStorage / 100.0)
}

type large struct {
	name string
	description string
	unitCost float64
}

func (l large) getPrice() float64 {
	return l.unitCost * (1 + largeStorage / 100.0) + shipping
}

func factory(size, name, description string, price float64) (product, error) {
	switch size {
	case smallSize:
		return small{name, description, price}, nil
	
	case mediumSize:
		return medium{name, description, price}, nil
	
	case largeSize:
		return large{name, description, price}, nil

	default:
		return nil, errors.New("ese tamaño no existe")
	}
}


func main() {
	peque, err := factory(smallSize, "Clip", " Clip de papeles", 10.5)

	if err == nil {
		fmt.Println("producto1: ", peque)
		fmt.Println("precio:", peque.getPrice())
	} else {
		fmt.Println(err.Error())
	}

	med, err1 := factory(mediumSize, "Calculadora", "Calculadora de mano", 1500.0)
	
	if err1 == nil {
		fmt.Println("producto2: ", med)
		fmt.Println("precio:", med.getPrice())
	} else {
		fmt.Println(err1.Error())
	}

	grande, err2 := factory(largeSize, "Monitor", "Monitor de PC", 35000.0)
	if err2 == nil {
		fmt.Println("producto3: ", grande)
		fmt.Println("precio:", grande.getPrice())
	} else {
		fmt.Println(err2.Error())
	}
}