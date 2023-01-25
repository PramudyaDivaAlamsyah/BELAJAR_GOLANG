package main

import "fmt"

func main(){
	fmt.Println("Ini String")
	fmt.Println(len("jumlah string"))
	fmt.Println("ambil karakter ke 0"[1]) //akan mengambil byte dari index ke 0 yaitu a, nanti akan dirubah lagi ke angka aslinya
	//bukan ke bytenya
}

//function string
// len("string") //menghitung jumlah karakter string
// "string"[number] //mengambil karakter pada posisi yang ditentukan