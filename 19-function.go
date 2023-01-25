package main

import (
	"fmt"
)

func sayMe(){
	fmt.Println("me")
}

//function parameter, harus dikasih type datanya
func sayParam(firstname string, lastname string){
	fmt.Println(firstname, " dan ", lastname);
}

//return functipn, harus membuat type kembalian di function tersebut
func getHello(name string) string {
	return name;
}

//return multiple function
func getMultiple(name string, lastname string) []string{
	arrays := [...]string{
		name, lastname,
	}

	return arrays[:];

}

//MULTIPLE RETURN
func getMultipleName()(string, string,string){//membuat returnnya dengan kurung dan type data(string, string)
	return "dipa","middle","alamsyah"
}

func main(){
	sayMe()
	sayMe()
	lastname := "alam"
	sayParam("dipa", lastname)

	name := getHello("dipa");
	fmt.Println(name)

	//panggil multiple return, dengan , dipisahkan, firstname untuk yang pertama, lastname untuk ynag kedua
	firstname, middlename, lastname := getMultipleName();
	fmt.Println(firstname,middlename,lastname)

	//mengabaikan salah satu returnnya gunakan _
	firstname2, _, lastname:= getMultipleName();
	fmt.Println(firstname2,lastname)



}