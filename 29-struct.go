package main

import "fmt"

//struct agar type data yang dimasukan berbeda2 sama halnya seperti class di java
type Customer struct{//struct adalah prototype data
	Name, Address string
	Age	int
}

func main(){
	//cara satu membuat struct
	var dipa Customer;
	dipa.Name = "Dipa";
	dipa.Address = "alamat";
	dipa.Age = 23;

	fmt.Println(dipa)

	//untuk print satu satu
	fmt.Println(dipa.Address);
	fmt.Println(dipa.Age);
	fmt.Println(dipa.Name)

	//cara dua membuat struct
	dipa2 := Customer{
		Name:"bebas",
		Address: "gatau",
		Age: 32,
	}

	fmt.Println(dipa2)
	fmt.Println(dipa2.Address);
	fmt.Println(dipa2.Age);
	fmt.Println(dipa2.Name)


	//cara 3 membuat struct
	dipa3 := Customer{"Dipa","Indonesia",30}//cetak sesuai nama di structnya dan posisinya harus sama

	fmt.Println(dipa3)
	fmt.Println(dipa3.Address)
	fmt.Println(dipa3.Age);
	fmt.Println(dipa3.Name)
}