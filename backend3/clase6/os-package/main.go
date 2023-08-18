package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	createFile("archivo3.txt")
	readFile("archivo2.txt")

	printDetails("archivo2.txt")
	os.Exit(0)
	fmt.Println("Esta linea no se va a mostrar")
}

func createFile(fileName string){
	 f, err := os.Create(fileName)
	 if err != nil {
		log.Fatal(err)
	 }

	 _, err2 := f.WriteString("Contenido del archivo")
	 if err2 != nil {
		log.Fatal(err)
	 }
}


func readFile(fileName string) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}

func printDetails(fileName string){
	f, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nEs un directorio:", f.IsDir())
	fmt.Println("Nombre del archivo/directorio:", f.Name())
	fmt.Println("Tamaño del archivo en bytes:", f.Size())
	fmt.Println("Fecha y hora del archivo:", f.ModTime())
	fmt.Println("Permisos del archivo:", f.Mode())
	
}