package main

import "fmt"

func main() {
    nums := []int{10, 5, 7, 3, 6}
    parCh := make(chan int)
    imparCh := make(chan int)

    go func() {
        for {
            select {
            case num := <-parCh:
                fmt.Printf("PAR: %d %s", num, "\n")
            case num := <-imparCh:
                fmt.Printf("IMPAR: %d %s", num, "\n")
            }
        }
    }()

    for _, num := range nums {
        if (num % 2) == 0 {
            parCh <- num
        } else {
            imparCh <- num
        }
    }
}