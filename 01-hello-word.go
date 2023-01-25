package main //untuk membuat program, packagenya harus main

import "fmt" //harus mengimport fmt ketika ingin mencetak

func main(){
	fmt.Println("Hello WOrd");//di golang bisa diakhiri dengan titik koma atau tidak
}

//go build 01-hello-word.go = untuk menjalankan filenya, masukan perintah go build <nama_file>.go di terminal
//eksekusi dengan ./01-hello-word di terminal hasil yang sudah di build tersebut

//go run <nama_file>.go untuk development


//ada beberpa tipe data number di go lang
//integer
/*

//integer ada minnya
int8
int16
int32
int64

//unsigned tak ada nilai minnya / dimulai dari 0
uint8
uint16
uint32
uint64

//float
float32
float64

//complex number biasa digunkana untuk matematiknya banyak
complex64
complex128

//alias = nama lain untuk tipe data yang sudah ada
byte === uint8
rune === int32
int === minimal int32 //tergantung sistem operasinya jika 32bit maka int32 jika 64bit maka int64
uint ===  minimal uint32 //tergantung sistem operasinya jika 32bit maka int32 jika 64bit maka uint64

//boolean
menggunakan keyword bool untuk merepresentasikan boolean

*/