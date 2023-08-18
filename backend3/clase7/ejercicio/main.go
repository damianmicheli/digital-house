package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main () {
	
	refundChannel := make(chan string)

	loanChannel := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			refundChannel <- "Refund Processed"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond*500)
			loanChannel <- "Loan Processed"
		}
	}()

	// **** version 1*****

	// go func(){
	// 	for {
	// 		fmt.Println(<- refundChannel)
	// 	}
	// }()

	// go func(){
	// 	for {
	// 		fmt.Println(<- loanChannel)
	// 	}
	// }()
	
	// fmt.Scanln()

	// *** fin version 1 ****

	// ** version 2 ***
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	
	for {
		select {
			
		case value := <-loanChannel:
			fmt.Println(value)
			
		case value:= <-refundChannel:
			fmt.Println(value)
			
		case <-interrupt:
			return
		}

	}

// **** fin version 2 ***
	
}