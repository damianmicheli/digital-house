package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)


type persona struct {
	Nombre string `json:"name"`
	Apellido string `json:"lastname"`
	Edad int
	Direccion string
	Telefono string
	EstaActivo bool
}

func main() {

	router := gin.Default()

	personaUno := persona{"Damian", "Micheli", 43, "Calle falsa 123", 
		"123456", false}


	personaByte, err := json.Marshal(personaUno)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(personaByte))

	contextGetFunc := func(context *gin.Context) {
		context.JSON(http.StatusOK, personaUno)
	}

	router.GET("/persona",contextGetFunc )

	router.Run(":8081")

}
