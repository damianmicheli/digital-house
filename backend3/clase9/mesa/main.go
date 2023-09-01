package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock int `json:"stock"`
	Codigo string `json:"codigo"`
	EstaPublicado bool `json:"estaPublicado"`
	FechaDeCreacion string `json:"fechaDeCreacion"`
}

type Productos struct {
	Productos []Producto `json:"productos"`
}


func main (){

	rawData, err := os.ReadFile("./productos.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Se abrio el archivo correctamente")

	var productos Productos

	err = json.Unmarshal(rawData, &productos)
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range productos.Productos{
		fmt.Println(p)
	}

	router := gin.Default()

	contextGetFunc := func(context *gin.Context) {
		context.JSON(http.StatusOK, productos.Productos)
	}

	router.GET("/productos",contextGetFunc )
	
	err = router.Run()
	if err != nil {
		fmt.Println(err)
	}

}