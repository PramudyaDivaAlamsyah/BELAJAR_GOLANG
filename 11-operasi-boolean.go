package main

import "fmt"

func main(){
	var nilai = 90;
	var absensi = 80;

	var lulusNilaiAkhir bool = 	nilai > 80;
	var lulusAbsensi bool = absensi > 80;

	lulus := lulusNilaiAkhir && lulusAbsensi;

	fmt.Println(lulus)
}