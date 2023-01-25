package main

import "fmt"

func getGoodBye(name string) string{
	return "good bye"
}

func main(){
	fuctionV := getGoodBye;

	fmt.Println(fuctionV("gogoy"))
}