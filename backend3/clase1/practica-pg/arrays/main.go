package main

import "fmt"

func main() {
	
	var b = [2]string {"mesa", "silla"} //array
	
	var a [2]string //arrsy
	
	var c = []string{"perro","gato","Loro", "hamster", "Loro", "hamster"} //slice

	a[0]="mesa"
	a[1]="silla"
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := make([]string, 5)
	d[3] = "arbol"
	fmt.Println(d)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(cap(c))
	fmt.Println(len(c))

	primos := []int {2,3,5,7,11}

	fmt.Println(primos[:3])

	// cap era 6, pero cuando agrego un loro pasa a 12
	c = append(c, "Loro")
	fmt.Println(cap(c))
	fmt.Println(len(c))
	fmt.Println(c)

	//declarar un map

	map1 := map[string]int{} // crear map vacio
	fmt.Println(map1)

	map2 := make (map[string]string) // crear map inicializado
	fmt.Println(map2)

	map3 := map[string]int{"d":3, "f":4} // crear map con contenido
	fmt.Println(map3)
	fmt.Println(len(map3)) // tamao del map
	fmt.Println(map3["f"]) // imprimir un elemento de un map

	map3["h"] = 25 // agregar un elemento al map
	fmt.Println(map3)

	map3["h"] = 255 // modificar un elemento del map
	fmt.Println(map3)

	delete(map3, "h") //borra elemento map
	fmt.Println(map3)

	// usar punteros
	var variable int = 45
	memoria := &variable
	fmt.Println(variable)
	fmt.Println(*memoria)



}