package main

import (
	"fmt"
	"time"
)

func main() {
	
	// go printNumber(2, nil)
	
	// go printNumber(3, nil) // con go le indico a la func que es concurrente (no se en que orden ejecutan)
	
	
	intChannel := make(chan int, 3) // canal que transporta datos de tipo entero. el 3 es un buffer de 3
	
	
	// unbuffered. no guarda informacion. buffered si
	
	go printNumber(intChannel) // no se ejecuta hasta que entre algo. queda escuchando
	intChannel <- 3 // 3 es un dato tipo entero que empieza a entrar en el canal
	intChannel <- 4 // 3 es un dato tipo entero que empieza a entrar en el canal
	intChannel <- 5 // 3 es un dato tipo entero que empieza a entrar en el canal
	time.Sleep(time.Second)

}

func printNumber(intChannel chan int) {
	
	value := <- intChannel
	
	fmt.Printf("Number is %d\n", value)

}