package main

import "fmt"

func main() {

	monthMap := make(map[int]string)

	monthMap[1] = "January"
	monthMap[2] = "February"
	monthMap[12] = "December"

	month := 2

	value, exists := monthMap[month]

	if exists {
		fmt.Printf("The month is %v", value)
	} else {
		fmt.Println("The month doesn't exists")
	}
}