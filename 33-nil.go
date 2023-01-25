//nil adalah data kosong, hanya bisa diberikan di beberapa type data seperti interface, function, map,
//slice, pointer dan channel

package main

import "fmt"

func NewMap(name string)map[string]string{
	if(name == ""){
		return nil
	}else{
		return map[string]string{
			"name":"duduy",
		}
	}
}

func main(){
	var person map[string]string; //jika tidak pakai nil makai if akan lebih panjang menyebutkan variablenya
	if(person["name"] == ""){
		fmt.Println("kosong")
	}else{
		fmt.Println(person)
	}
	

	person2 := NewMap("awdwa");
	if(person2 == nil){
		fmt.Println("nil")
	}else{
		fmt.Println(person2)
	}
}