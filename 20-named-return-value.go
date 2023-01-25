package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string){
	firstname = "pramudya";
	middlename = "diva";
	lastname = "alamsyah";

	return; 
	//tidak perlu memanggil returnnya firstname, middlename, lastname lagi, karena sudah di deklarasikan diatas
	//func getC()(firstname, middlename, lastname); akan mereturn didalam ()(sini)
}

func main(){
	a, middlename, lastname := getCompleteName();//nama gak harus sama dengan yang di return
	fmt.Println(a, middlename, lastname)
}