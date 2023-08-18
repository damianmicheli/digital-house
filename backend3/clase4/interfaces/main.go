package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

const (
	rectType = "rect"
	circleType = "circle"
)

// FIGURA GEOMETRICA

type geometry interface {
	area() float64
	perim() float64
}

func newGeometry(geoType string, values ...float64) geometry {

	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}

	return nil
}

func details(g geometry) {
	fmt.Printf("Tipo: %s\n", getType(g))
	fmt.Println("Valores:", g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimetro: %.2f\n", g.perim())
}

func getType(g geometry) string {

	fullTypeName := reflect.TypeOf(g).String()
	sliceFullTypeName := strings.Split(fullTypeName, ".")
	typeName := sliceFullTypeName[len(sliceFullTypeName)-1]

	switch typeName {
		case rectType:
			typeName = "Rectángulo"
		case circleType:
			typeName = "Círculo" 
	}

	return typeName
}


// CIRCULO

type circle struct {
	radius float64
}

func newCircle(values float64) circle {
	return circle{radius: values}
}


func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


// REACTANGULO

type rect struct {
	width, height float64
}
func newRect(w, h float64) rect {
	return rect{width: w, height: h}
}
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}


// INTERFAZ VACIA

type ListaHeterogenea struct {
	Data []interface{}
}

// MAIN

func main() {

	c := newCircle(5)
	c2 := newGeometry(circleType, 6)
	r := newRect(16, 9)
	r2 := newGeometry(rectType, 4, 3)
	details(c)
	details(c2)
	details(r)
	details(r2)

	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "hola")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64)
	// fmt.Println(f)
}