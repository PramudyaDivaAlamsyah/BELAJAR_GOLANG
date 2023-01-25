package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string;

func main(){
	parameterFunc("Anjing", spamFilter);

}

func spamFilter(name string) string{
	if name == "Anjing" {
		return "...";
	}else{
		return name;
	}
}

func upperFilter(name string) string{
	dipaUpper := strings.ToUpper(name);
	return dipaUpper
}

//function parameter
func parameterFunc(name string, filter Filter){
	nameFiltered := filter(name);

	fmt.Println(nameFiltered);

	fmt.Println(upperFilter("dipa"))
}

//type declaration, membuat alias pada sebuah function jika parameter type banyak