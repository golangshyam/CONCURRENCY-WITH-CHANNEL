package main

import (
	"fmt"
)

func receiver(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func generator() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
func main() {
	c := generator()
	receiver(c)
}
