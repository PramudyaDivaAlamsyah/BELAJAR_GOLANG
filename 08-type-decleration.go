package main

import "fmt"

func main(){
	type NoHp string; //type untuk mendifinisikan kalo NoHp adalah sebuah type yang typenya string;
	type Married bool;

	var noHp NoHp; //variabel noHp yang bertipe NoHp
	noHp = "dipa"

	var married Married = true;


	fmt.Println(noHp)
	fmt.Println(married)
}