package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code int
}

func main() {
	err := doSomething()
	var customErr *MyError

	if errors.As(err, &customErr) {
		fmt.Printf("Código de error personalizado: %d\n", customErr.Code)
	} else {
		fmt.Println("No se pudo extraer el error personalizado")
	}
}

func doSomething() error {
	// Simulating an error
	return &MyError{Code: 404}
}