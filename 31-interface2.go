package main

import "fmt"

func SayHello(tes Tes){
	fmt.Println("hello", tes.GetName())
}

type Tes interface{
	GetName() string
}


type Animal struct{
	Name string	
}

func (animal Animal) GetName() string {
	return animal.Name;
}

func main(){
	animal := Animal{
		Name: "bebaslah ya",
	}

	SayHello(animal);

}