package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int

}

func sayHi(customer Customer, name string){
	fmt.Println(customer.Name)
	fmt.Println(name)
}

//FUNCTION STRUCT
func (customer Customer) sayHi2(name string){
	fmt.Println(customer.Name)
	fmt.Println(name)
}

func main(){
	//cara tidak memakai struct method
	dipa := Customer{
		Name: "apa gitu",
		Address: "gatau",
		Age:300,
	}

	sayHi(dipa, "dipa nih")

	//cara pakai struct method
	dipa.sayHi2("kepo")
}