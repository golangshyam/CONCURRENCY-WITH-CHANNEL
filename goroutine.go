package main

import (
	"fmt"
)
/*
func printIntegers(done <-chan struct{}, intStream <-chan int) {

	fmt.Println("printIntegers goroutine ")

	for {

		fmt.Println("infinite  loop ")
		select {
		case i := <-intStream:
			fmt.Println("int data existed")
			fmt.Println(i)
		case <-done:
			return
		}
	}

	fmt.Println("end of infinite loop ")

}

func main() {

	var chance chan struct{} = make(chan struct{})
	var chance1 chan int = make(chan int)
	go printIntegers(chance, chance1)
	chance1 <- 411
	fmt.Println("main existed")

}


func printIntegers(done <-chan struct{}, intStream <-chan int) {

	fmt.Println("printIntegers goroutine ")

	for {

		fmt.Println("infinite  loop ")
		select {
		case i := <-intStream:
			fmt.Println("int data existed")
			fmt.Println(i)
		case <-done:
			fmt.Println("struct data existed")
			return
		}
	}

	fmt.Println("end of infinite loop ")

}

func main() {

	var chance chan struct{} = make(chan struct{})
	var chance1 chan int = make(chan int)
	go printIntegers(chance, chance1)
	type coder struct {
		Name string
	}
	var cod coder=coder{"shyam"}
	fmt.Println(cod)
	ret:=coder{}
	chance <- ret
	fmt.Println("main existed")

}
*/
// define a struct type
func printstructure(ch <- chan Person, ch1<- chan shyam()){
	fmt.Println("Received:", <-ch)
	for {
		fmt.Println("loop will excute infinite times")
		select{
		case i:=<-ch:
			fmt.Println("person structute data is:",i)

		case <- ch1:
			fmt.Println("print the shyam function data")
		}
	}
}
func shyam (){
	var Name string="sam"
	var Age int=24
	fmt.Sprintln(Name,Age)
}
type Person struct {
    Name string
    Age  int
}

func main() {
    // create a channel of type Person
    ch := make(chan Person)
	ch1:= make(chan shyam())
	go printstructure(ch,ch1)

    // start a goroutine to receive data from the channel
   // go func() {
        // receive data from the channel
   //     p := <-ch
        // print the received data
   //     fmt.Println("Received:", p)
   // }()

    // create a Person struct and send it to the channel
	ret:=shyam()
    p := Person{Name: "shyam", Age: 24}
    ch <- p
	ch1<-ret
    // wait for user input before exiting the program
    fmt.Scanln()
}
package main

import "fmt"

// define an interface with multiple methods
type Shape interface {
    Area() float64
    Perimeter() float64
}

// define a struct for a rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// implement the Shape interface for the Rectangle struct
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

func main() {
    // create a Rectangle object
    rect := Rectangle{Width: 10, Height: 5}

    // call the methods of the Rectangle object
    fmt.Println("Area:", rect.Area())
    fmt.Println("Perimeter:", rect.Perimeter())

    // create a Shape interface variable and assign the Rectangle object to it
    var s Shape
    s = rect

    // call the methods of the Shape interface
    fmt.Println("Area:", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}







