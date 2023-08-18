package main

import "fmt"

func main() {

	p, q := 58, 22

	if p > q { 
		fmt.Println(p, "es mayor que", q)
	} else if p < q {
		fmt.Println(p, "es menor que", q)
	} else {
		fmt.Println(p, "y",  q, "son iguales" )
	}

	
	if r, s := 255, 22; r > s && p < q && "a" > "A"{ 
		fmt.Println(r, "es mayor que", s, "y", p, "es menor a", q )
	}



	opcion := "compefrar"

	switch opcion {

		case "comprar": 
		fmt.Println("pero vo quere comprar")
		fmt.Println("y aca te sigo hablando")
		
		case "vender": 
		fmt.Println("pero vo quere vender")
		
		default: 
		fmt.Println("pero yo no se que es lo q quere vo")

	}

	//sin condicion
	switch { //tiro condiciones y va la primera que se cumple
	case p == q:
		fmt.Println(p, "es mayor que", q)
		
	case 8 == 9:
		fmt.Println("8 es menor que nueve")
		
	default:
			fmt.Println("aca no se cumple nada")
	}

	//con multiples casos

	expresion := "blanco"

	switch expresion {
		case "rojo", "azul", "naranja":
			fmt.Println("color del primer grupo")
		case "verde", "amarillo", "violeta":
			fmt.Println("color del segundo grupo")
		default:
			fmt.Println("color de ningun grupo")
		}


	// declaracion corta y fallthrough (lo opuesto de break)
	switch expresion1 := "rojo"; expresion1 {
	case "rojo", "azul", "naranja":
	
		fmt.Println("color del primer grupo")
	
		fallthrough
	
	case "verde", "amarillo", "violeta":
		
		fmt.Println("color del segundo grupo")
	
		if expresion1 != "azul" {
	
			break //interrupe el flujo
	
		}
	
		fmt.Println("color del segundo grupo otra vez")

	default:
		fmt.Println("color de ningun grupo")
	}


}