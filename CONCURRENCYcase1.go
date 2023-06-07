package main

import (
	"fmt"
)

func prints(done <-chan struct{}, Intstream <-chan int) {

	fmt.Println("prints go routine started")

	for {
		fmt.Println("infinite loop started")
		select {
		case i := <-Intstream:
			fmt.Println("int data existed")
			fmt.Println(i)
		case <-done:
			return
		}
	}

	
}

func main() {

	var chance chan struct{} = make(chan struct{})

	var chance1 chan int = make(chan int)

	go prints(chance, chance1)

	chance1 <- 1

	fmt.Println("main ends")
}
