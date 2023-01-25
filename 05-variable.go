package main

import "fmt"

func main(){
	var namaVariabel string
	//tidak bisa di isi dengan tipe data variable lain

	namaVariabel = "ini variabel tipe data string"
	fmt.Println(namaVariabel);

	namaVariabel = "terganti"
	fmt.Println(namaVariabel)

	var nama = "bisa mengisi dengan tipe data seperti valuenya";
	fmt.Println(nama);

	var age = 30; //tipe datanya akan int minimal 32 tergantung sistem operasi yang digunakan
	//atau bisa dipaksa
	var age2 int8 = 30; //tipe datanya jadi int8
	fmt.Println(age);
	fmt.Println(age2);

	// var age; //tak bisa seprti init jika ingin di inisiasikan harus dikasih tipe datanya

	//cara agar tidak menggunakan var
	country := "tanpa var"; //:= untuk deklarasi awal
	fmt.Println(country);

	// country := "jangan begini" //tidak boleh begini untuk mengganti datanya
	country = "hehe" //tapi begini

	//deklarasi multiple variable
	var(
		firstname = "dipa";
		lastname string
	)

	//jika digolang variable harus dipakai, tidak bisa jika tidak dipakai

	fmt.Println(firstname);
	fmt.Println(lastname)

}