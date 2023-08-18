package main

import "fmt"

func main() {

	nombres := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Antes:")
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	nombres = append(nombres, "Gabriela")
	fmt.Println("Después:")
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	count := 0
	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])
	for _, age := range employees {

		if age > 21 {
			count++
		}
	}
	fmt.Println(count)

	employees["Federico"] = 25
	for name := range employees {
		fmt.Println(name)
	}

	delete(employees, "Pedro")
	for name := range employees {
		fmt.Println(name)
	}

}
