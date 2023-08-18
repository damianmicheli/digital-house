package main

import "os"

func main (){

	categorias := "./categorias.csv"

	file, err := os.Create(categorias)
	if err != nil {
		panic(err)
	}

	defer func() {
		file.Close()
	}()

	const ContenidoArchivo = `001,Electrodomésticos,TV-Heladera
002,Muebles,Mesa-Silla
003,Herramientas,Martillo-Destornillador
004,Pinturas,Latex-Sintetico
005,Aberturas,Puerta-Ventana
006,Construcción,Ladrillo-Cemento
007,Automotor,Moto-Auto`

	_, err = file.WriteString(ContenidoArchivo)
	if err != nil {
		panic(err)
	}

}

func cargarDatos(){
	
}