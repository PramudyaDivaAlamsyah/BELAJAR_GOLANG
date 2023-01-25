package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int)(int, error){
	if pembagi == 0 { //membuat errors
		return 0, errors.New("Pembagi tidak boleh 0")
	}else{
		return nilai/pembagi, nil; //nil jika tidak terjadi error
	}
}

func main(){
	var contohError error = errors.New("Ada Error Nih");//bawaan dari golang untuk melempar error
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(1,5);

	if err == nil{
		fmt.Println("Hasil ", hasil);
	}else{
		fmt.Println("Error", err.Error())
	}

	fmt.Println(hasil);
}	