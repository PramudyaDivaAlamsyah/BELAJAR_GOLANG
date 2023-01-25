package main

import "fmt"

func main(){
	//tipe data slice adalah potongan dari array, ukuran slice bisa berubah
	//tipe data slice memiliki	3 data, yaitu pointer, length, capacity

	//pointer = penunjuk data pertama di array pada slice
	//length  = panjang dari slice
	//capacity = kapasitas dari slice, dimana length tidak boleh lebih dari capacity
	//capacity berjumlah awal dari pointer ke paling akhir index array tersebut

	//cara membuat slice
	//array[low:hight] //membuat slice dari array dimulai index low sampai index sebelum hight
	//array[low:] //membuat slice dari array dimulai index ke low sampai index akhir di array
	//array[:hight] //membuat slice dari array index 0 sampai index sebelum hight
	//array[:] // membuat slice dari array dimulai index 0 sampai index akhir di array

	month := [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Oktober",
		"September",
		"November",
		"Desember",
	}
	fmt.Println(month);

	//[pointer:length]

	slice := month[1:4]; //ambil slice dari index ke 1 sampai index sebelum 3 // akan memakai reference, 
	//jika diganti maka akan ikut terganti semuanya

	//len untuk mengambil panjangnya
	fmt.Println(len(slice))
	
	//cap untuk mengambil capacitynya
	fmt.Println(cap(slice)); //capacity didapat dari pointer dihitung ke posisi paling akhir index array

	slice[0] = "berubah"; //ini akan menunjuk array ke 0pada slicenya
	fmt.Println(slice);
	fmt.Println(month)

	//append untuk membuat slice baru dengan menambah ke terakhir slicenya (akan menambah data baru/tidak reference jika
	//arraynya sudah penuhh)
	var slice2 = append(slice,"dodol");

	fmt.Println(slice2);
	fmt.Println("yp",month); //akan terganti

	//akan dibuat array baru jika melebihi kapasitas arraynya
	var slice3 = month[10:];
	var slice4 = append(slice3,"katak");

	fmt.Println(slice4)
	fmt.Println(month) //tidak akan terganti karena dia akan membuat refenrece atau alamat baru

	//make untuk membuat slice baru (yang paling aman membuat slice karena tidak ada variablenya/var)
	newSlice := make([]string,2,7); //buat slice yang tipenya array string, panjangnya 2 dan capacitynya 7
	newSlice[0] = "dipa";
	newSlice[1] = "alamsyah";
	// newSlice[2] = "aw";//akan error karena panjangnya hanya 2

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice));//capacitynya 7
	fmt.Println(len(newSlice));//lengthnya 2

	//copy slice untuk mengcopy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice);//copy slice, parameter 1 tempat buat menaruh copyan, parameter 2 yang mau di copy
	fmt.Println("copy",copySlice)


	//hati2 saat membuat array
	iniArray := [...]int{1,2,3,4,5,6}
	iniSlice := []int8{1,2,3,4,5,6}

	fmt.Println(iniArray);
	fmt.Println(iniSlice)




}