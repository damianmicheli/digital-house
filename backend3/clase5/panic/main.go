package main

import "fmt"

func main () {
num := 3
isPair(num)
fmt.Println("Ejecucion completada!")
}

func isPair(num int){

	func(){
		fmt.Println("hola")
	}()

	defer func(){
		err := recover()
		fmt.Println("Aca entra siempre")

		if err != nil {
			fmt.Println(err)
		}
	}()

	if(num % 2) != 0 {
		panic(fmt.Sprintf("%d no es un numero par",num))
	}

	fmt.Println(num, "es un número par")


}


