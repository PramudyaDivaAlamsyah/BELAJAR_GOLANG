package main

import "fmt"

func main() {
	name := "dip"

	//di golang if expression tanpa kurung buka dan kurung tutup
	if name == "dipa" {
		fmt.Println("clear")
	}else{
		fmt.Println("no way")
	}

	//if sort statement hanya ada di golang
	//name:="dip"; adalah sort statement yang mana kita mendeklarasikan value yang ingin di if di dalam if
	if name:=len("dip"); name==3{
		fmt.Println("bener cuy")
	}else{
		fmt.Println("salah cuy")
	}
}