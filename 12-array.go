package main

import "fmt"

func main(){//cara membuat array, di golang tipe data array hanya bisa satu tidak bisa bermacam2

	//membuat array manual
	var names [3]string;//arrray dengan tipe data string
	names[0] = "dipa";
	names[1] = "aku";
	names[2] = "bebas";

	// //error
	// names[3] = "ahay";// karena kelebihan... hanya menirima assign array sampai 3 size saja
	// names[2] = 3; //error karena hanya menerima tipe data array yang tipenya string

	fmt.Println(names)

	//membuat array langsung
	var values = [...]int{ //bisa bebas seperti ini tanpa menentukan maksnya
		90,91,92,12,12,12,
	}
	fmt.Println(values)

	var values2 = [3]int{ //atau ditentukan maks arraynya
		90,91,92,
	}

	fmt.Println(values2)

	//function array untuk memanipulasi array
	var jumlahArray = len(values);//cek panjang dari arraynya bukan dari isinya
	var jumlahArray2 int8 = int8(len(values2)); //bisa di convert

	fmt.Println(jumlahArray)
	fmt.Println(jumlahArray2)
}