package main

import "fmt"

func main(){
	counter := 1;

	//perulangan di go lang hanya ada for
	for counter <= 10{ //bisa ddibuat seperti ini
		fmt.Println(counter);
		counter++;
	}


	//atau dengan statement seperti ini
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println(counter2)
	}

	//for range untuk melakukan perulangan di data array, slice, atau map
	slice := []string{"DIpa","gans","alam"};

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])	
	}


	for index, value := range slice {  //range lalu masukan array yang ingin di cetak
		fmt.Println(value, " index ke", index)
	}

	for _ ,value := range slice {  //jika index tak ingin dipakai, maka ganti dengan _ (underscore)
		fmt.Println(value, " index ke",)
	}


	person := make(map[string]string);

	pakai := map[string]string{
		"apa":"we",
	}

	fmt.Println(pakai)

	person["name"] = "dipa";
	person["las"] = "alam";

	for key, value := range person {
		fmt.Println("key", key);
		fmt.Println("value", value)
	}

}