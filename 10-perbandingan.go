package main

import "fmt"

func main(){
	var name1 = "dipa";
	var name2 = "cipa";

	var banding bool = name1 == name2;

	fmt.Println(banding);

	var compare bool = name1 > name2; //bisa di compare dari hurufnya, d lebih besar dari c maka true

	fmt.Println(compare)

	tes := "aku";

	fmt.Println(tes);
}