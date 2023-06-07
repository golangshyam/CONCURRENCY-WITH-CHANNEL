package main

import(
	"fmt"
	_"time"
)
/*
func foo(s string){
	for i:=1; i<=3; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s,i)
	}
}
func main(){
	fmt.Println("main go routine started")

	go foo("1st goroutine")
	go foo("2nd goroutine")

	time.Sleep(time.Second)
	fmt.Println("main fo routine ended")
}

// write and read from channel
func shyama(shy chan int){
	for i:=0; i <5; i++{
		shy<-i
	}
}
func main(){

	shy := make(chan int)

	go shyama(shy)

	for i:=0; i < 5; i++{
		fmt.Println(<-shy)
	}
} // 0 1 2 3 4

// channel close statement and excusion

func shyama(shy chan int){
	for i:=0; i <5; i++{
		shy<-i
	}
	close(shy)
}
func main(){

	shy := make(chan int)

	go shyama(shy)

	for i:=0; i < 6; i++{
		fmt.Println(<-shy)
	}
}//0 1 2 3 4 5 0 the last one zero is printing without blocking iteration
// means channel closed
*/
// channel default and declration

func main() {  
    var c chan int
    if c == nil {
        fmt.Println("channel c is nil, going to define it")
        c = make(chan int)
        fmt.Printf("Type of c is %T", c)
    }
}

