// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	_"time"
)
/*
type microservices struct {
	name  string
	date  string
	price int
}

func server(data chan microservices) {

	func() {

		var transform map[string]interface{} = make(map[string]interface{})

		for {

			select {

			case datas := <-data:

				if datas.name != "" && datas.date != "" && datas.price != 0 {

					transform["name"] = datas.name
					transform["date"] = datas.date

					fmt.Println("transformed data ", transform)

				}

			}
			fmt.Println("server go routines")

		}

	}()

}

func client(data chan microservices) {

	fmt.Println("client")

	data <- microservices{"goa", "march 2020", 16000}

	fmt.Println("client  completed ")

}

func main() {

	var data chan microservices = make(chan microservices)

	fmt.Println("channel creation ")

	go client(data)

	fmt.Println("gorouine registration ")

	server(data)
	time.Sleep(1)

}
*/
type microservice struct{
	name string
	date string
	price int
}

func client (data chan microservice){
	fmt.Println("client registred")
	data<-microservice{"shyam","20dec2023",20000}
	fmt.Println("client go registred")
}
func server (data chan microservice){
	fmt.Println("server go routine registred")
	func(){
	var shyam map[string]interface{}=make(map[string]interface{})
	for {
		select{
		case datas:=<-data:
			if datas.name!=""&& datas.date!=""&&datas.price!=0{
				shyam["name"]=datas.name
				shyam["date"]=datas.date
				shyam["price"]=datas.price
				fmt.Println(shyam)

			}
		}
		fmt.Println("server registred")
	}

	


	}()


}

func main(){

	var data chan microservice=make(chan microservice)
	go client(data)
	server(data)
	fmt.Println("server client read data completed")
}