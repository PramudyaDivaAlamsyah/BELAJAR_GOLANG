package main

import "fmt"

type BlackList func(string) bool;

var blackList = func(name string) bool {//harus pakai var jika mau global
	return name == "admin";
}

func registerUser(name string, blackList BlackList){
	if(blackList(name)){
		fmt.Println("anda di blokir",(name));
	}else{
		fmt.Println("welcome", name)
	}
}

func main(){
	blackList2 := func(name string) bool {//harus pakai var jika mau global
		return name == "admin";
	}

	registerUser("admin",blackList2)
	
	registerUser("admic", blackList)
	registerUser("root", func(name string) bool{
		return name == "root";
	})
}