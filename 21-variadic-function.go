package main

import "fmt"

//namanya variadic function
func sum(numbers ...int) int{//hanya bisa divariable belakang, tidak bisa di awal atau tengah
	nilai := 0;

	for _, data := range numbers {
		nilai += data;
	}
	return nilai;
}

func main(){
	numbers := sum(10,20,30,40,50);
	fmt.Println(numbers)

	//slice parameter
	numcuys := []int{10,10,10,10};//ini slice
	numbers = sum(numcuys...);//akan dipecah slicenya, agar bisa di sum, hanya bisa untuk slice

	fmt.Println(numbers)
}