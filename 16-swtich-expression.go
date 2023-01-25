package main

import "fmt"

func main() {

	var name string = "dipdwol"

	switch name {
	case "bobol":
		fmt.Println("huyesa")
	case "dx":
		fmt.Println("salah")
	default:
		fmt.Println("defalt")
	}

	//switch dengan sort statement
	switch length:=len(name); length>5{
	case true:
		fmt.Println("lebih besar dari 5")
	case false:
		fmt.Println("lebih kecil dari 5")
	}

	//switch condition
	length2:=len(name);

	//tidak usah di panggil di switchnya length2 tersebut langsung di case
	switch { //mirip dengan if else
	case length2 > 2:
		fmt.Println("lebih besar dari 2")
	case length2 > 5:
		fmt.Println("nama terllau panjang")
	default:
		fmt.Println("nama sudah benar")
	}

}