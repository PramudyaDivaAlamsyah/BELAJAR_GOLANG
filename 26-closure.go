package main

import "fmt"

func main(){
	counter := 0;
	increment := func ()  {//berinteraksi dengan data yang ada di scope tersebt
		fmt.Println("Increment");
		counter++;//counter bisa di akses dari atas ke bawah tapi tidak bisa diakses dari bawah ke atas
	}

	nama := func(){
		counter := 1;//jika dikasih := atau deklarasi variable.. variable diatasnya tidak akan terganti
		counter++;
		fmt.Println(counter);
	}

	increment();
	increment();
	nama();
	fmt.Println(counter)
}