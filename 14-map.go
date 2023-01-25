package main

import "fmt"

func main(){
	//map terdiri dari key dan value map[typeKey]typeValue
	person := map[string]string{
		"name":"dipa",
		"address":"tangerang",
	}

	fmt.Println(person);


	//function map
	//len mengambil jumlah data yang ada di map
	fmt.Println(len(person))//2

	//map[key] //mengambil map lewat key
	fmt.Println(person["name"]); //dipa
	fmt.Println(person["address"]); //tangerang

	//map[key] = value //mengganti atau mengisi map dengan value sesuai key
	person["name"] = "berubah";
	fmt.Println(person);

	//make(map[typeKey]typeValue) //untuk membuat map baru
	var data map[string]string = make(map[string]string);
	data["name"] = "jojo";
	data["action"] = "berak";

	fmt.Println(data)

	//bisa gini tanpa memasukan tipe data di var
	var data2 = make(map[string]string);

	data2["name"] = "joxjo";
	data2["action"] = "berxak";

	fmt.Println(data2)

	//delete(map,key) // untuk menghapus map lewat key
	
	//sebelum dihapus 2 lennya
	delete(data2, "name"); // menghapus key name

	//setelah dihapus jadi 1 len nya
	fmt.Println(data2);
	


}