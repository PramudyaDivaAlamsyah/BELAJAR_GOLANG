//interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
//contohnya
/** fmt.Println
panic
recover
dan lain2
*/

package main

import "fmt"

//interface kosong
//bisa mendefinisikan parameter dengan type data apapun dengan interface kosong
func Ups(i int, apapun interface{}) interface{}{
	if i == 1 {
		return 1
	}else if i == 2{
		return true
	}else{
		return "Ups"
	}
}

//sama seperti
type Ups2 interface{};

func main(){
	interfaceKosong := Ups(1);//interface ksong akan menerima type data apapun
	fmt.Println(interfaceKosong);

	var interfaceKosong2 = Ups(2);
	fmt.Println(interfaceKosong2);

	var interfaceKosong3 interface{} = Ups(1)
	fmt.Println(interfaceKosong3)
}