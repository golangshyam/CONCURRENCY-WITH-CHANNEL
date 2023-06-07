package main

import (
	"context"
	"fmt"
)

func shyam(ctx context.Context, Intstream <-chan int) {

	for {
		select {
		case i := <-Intstream:
			fmt.Println(i)
		case <-ctx.Done():
			fmt.Println("context done")
			return
		}
	}
}

func main() {

	ctx := context.TODO()

	var comm chan int = make(chan int)

	go shyam(ctx, comm)

	comm <- 1

	fmt.Println("context interface and main ends")
}
