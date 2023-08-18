package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main () {
	error_1 := fmt.Errorf("primer error")
	error_2 := fmt.Errorf("segundo error: %w", error_1)
	error_3 := fmt.Errorf("tercer error: %w", error_2)

	fmt.Println(errors.Unwrap(error_3).Error())
	fmt.Println(error_3.Error())

	_, err := os.Open("not_exist.txt")
	// var notExist error = fs.ErrNotExist
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("The file does not exist")
	}

	
	_, err2 := os.Open("not_exist.txt")
	var pathError *fs.PathError


	if errors.As(err2, &pathError) {
		fmt.Println("Path error", pathError.Path)
	}


}