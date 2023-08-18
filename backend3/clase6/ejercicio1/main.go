package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	path := "./data.csv"

	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	writePath := "./datawrite.txt"

	file, err := os.Create(writePath)


	defer func() {
		file.Close()
	}()

	if err != nil {
		panic(err)
	}

	splittedRawDAta := strings.Split(string(rawData),"\n")

	var totalAmount float64

	for i, item := range splittedRawDAta {

		splittedItem := strings.Split(item,",")

		if i != 0 {

			price, err := strconv.ParseFloat(splittedItem[1], 64)

			if err != nil {
				panic(err)
			}

			quantity, err := strconv.ParseInt(splittedItem[2], 10, 64)
			
			if err != nil {
				panic(err)
			}

			itemAmount := price * float64(quantity)

			totalAmount += itemAmount


		}

		for j, item := range splittedItem {

			// fmt.Printf("%s\t\t", item)
			fmt.Fprintf(os.Stdout, "%s\t\t", item)
			fmt.Fprintf(file, "%s\t\t", item)
			
			if j == len(splittedItem)-1 {
				// fmt.Print("\n")
				fmt.Fprintf(os.Stdout,"\n")
				fmt.Fprintf(file,"\n")
			}
		}
		
	}

	// fmt.Printf("\n\nTotal:\t\t%.2f", totalAmount)
	fmt.Fprintf(os.Stdout,"\n\nTotal:\t\t%.2f", totalAmount)
	fmt.Fprintf(file,"\n\nTotal:\t\t%.2f", totalAmount)

}




