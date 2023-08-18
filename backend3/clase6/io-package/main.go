package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var s = "Este texto es copiado"

	r := strings.NewReader(s)
	_, err := io.Copy(os.Stdout, r)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(b))


	var s2 = "Este texto es copiado hasta la ultima linea"

	_, err2 := io.WriteString(os.Stdout, s2)

	if err2 != nil {
		log.Fatal(err2)
	}

}