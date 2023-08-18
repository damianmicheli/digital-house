package main

import "fmt"

type Vehiculo struct {
	km     float64
	tiempo float64
}

func (v Vehiculo) detalle() {
	fmt.Printf("km:\t%.2f\ntiempo:\t%.2f\n", v.km, v.tiempo)
}

type Auto struct {
	v Vehiculo
}

func (a *Auto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

type Moto struct {
	v Vehiculo
}

func (a *Moto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 80
}

func (a *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	a.v.detalle()
}

func main () {
	
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()

}