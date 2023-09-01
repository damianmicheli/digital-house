package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

const (
    categoryCSVPath string = "./files/categorias.csv"
)

func C6S()  {

    file := createCategoryFile(categoryCSVPath)
    defer file.Close()
    populateCategoryFile(file)

    rawData, err := os.ReadFile(categoryCSVPath)
    if err != nil {
        log.Fatal(err.Error())
    }

    splitRow := strings.Split(string(rawData), "\n")

    // estimaciones := map[string]float64{}

    for _, row := range splitRow {
        fmt.Println(row)
    }
    // recorrerer las categorias de el archivo categoria y por cada una recorrer el archivo de productos
    // crear un map contador con las categorias como propiedades y con el valor como valor.
    // si hay coincidencia calcular la cantidad por el valor y agregarlo a un contador de cada categoria.

    // Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.
    // Imprimir todos los estimativos por consola


    fmt.Fprint(os.Stdout, "Corriendo ejercicio C6S")
}

func createCategoryFile(path string) *os.File {
    file, err := os.Create(path)

    if err != nil {
        log.Fatal(err.Error())
    }

    return file
}

func populateCategoryFile(file *os.File) {

    const s string = `001,Electrodomésticos,tv-heladera
002,Muebles,silla-mesa
003,Herramientas,martillo-taladro
004,Pinturas,latex-sintetica
005,Aberturas,ventana-puerta
006,Construcción,ladrillo-cemento
007,Automotor,moto-auto`

    _, err := io.WriteString(file, s)
    if err != nil {
        log.Fatal(err)
    }

}