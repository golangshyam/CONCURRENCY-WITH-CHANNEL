package main

import (
	"fmt"
)

func print(done <-chan struct{}, Intstream <-chan int) {

	fmt.Println("print go routine started")

	for {
		fmt.Println("infinite loop started")

		select {
		case i := <-Intstream:
			fmt.Println("int dat existed")
			fmt.Println(i)

		case <-done:
			fmt.Println("struct data existed")

			return

		}
	}
	fmt.Println("gjjdl")
}

func main() {

	var ch chan struct{} = make(chan struct{})

	var ch1 chan int = make(chan int)

	go print(ch, ch1)

	type shyam struct {
	}
	ch <- shyam{}

	fmt.Println("main ends")
}
