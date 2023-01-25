package main

import "fmt"


func random() interface{} {
	return "dipa";
}

func main(){
	var result interface{} = random();
	var resultString string = result.(string);//type assertion, akan mengeluarkan string
	fmt.Println(resultString)

	//jika salah melakukan set assertion maka akan terjad panic
	//contoh
	// resultInt := result.(int);
	// fmt.Println(resultInt);//akan panic / error


	//type assertion with switch
	//untuk mengatasi panic agar program tidak terhenti saat panic
	//mengambil type dari value result

	switch value := result.(type){
	case string:
		fmt.Println("Ini string", value)

	case int:
		fmt.Println("ini int", value)

	default:
		fmt.Println("Unknown", value)
	}	
}