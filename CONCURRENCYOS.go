package main

import (
	"fmt"
	"time"
)

func Airhost(c chan string) {

	fmt.Println("terminal1")

	for i := 0; i < 10; i++ {

		fmt.Println("Looping started  ", i)

		select {

		case c <- "depature closed":
			fmt.Println("written in to channel")

		case <-time.After(1 * time.Second):

			fmt.Println("no waiting i am leaving ")
			break
		}

	}

	fmt.Println("Airhost before done")

	fmt.Println("Airhost after done")

}

func Gate(c chan string) {

	fmt.Println("gate before done")

	for i := range c {

		fmt.Println("terminal2 reading", i)
	}

	fmt.Println("gate after done")
}

func main() {

	var comm chan string = make(chan string) // heap memory
	// write or read os is informed

	go Airhost(comm)

	go Gate(comm)

	fmt.Println("two goroutines registered ")

	time.Sleep(1)

	fmt.Println("main ends  ")

}
