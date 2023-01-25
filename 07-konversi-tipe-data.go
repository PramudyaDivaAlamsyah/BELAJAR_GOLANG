package main

import "fmt"

func main(){
	//konversi tipe data
	var nilai32 int32 = 100000;
	var nilai64 int64 = int64(nilai32);

	var nilai16 int16 = int16(nilai32); //hati hati conversi ke tipe data yang lebih kecil,
	//jika datanya lebihbesar maka nilainya akan berbeda, akan kembali dari 0 lagi sampai cukup

	fmt.Println(nilai32);
	fmt.Println(nilai64);
	fmt.Println(nilai16);


	var nama = "ini dipa";

	//konversi byte to string
	var e = nama[0];//akan menyimpan byte dari string variable nama index ke 0
	var stringNama string = string(e);//ubah ke string

	fmt.Println(stringNama);
}