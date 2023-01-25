//defer adalah function yang dijadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
package main

import "fmt"

func logging(){
	fmt.Println("SELESAI MEMANGGIL FUNCTION")
}

func runApplication(value int){
	//akan dieksekusi meskipun terjadi error
	defer logging();// biasakan gunakan defer di paling atas
	fmt.Println("Run Application")
	result := 12/value;

	//jika terjadi error yang dibawah tidak akan di eksekusi (tanpa defer)
	fmt.Println("Result", result);
	logging();//tanpa defer

	runApp(true)
}

//panic function untuk menghentikan program jika terjadi error saat program berjalan
func runApp(error bool){
	defer endApp();
	if error {
		panic("Aplikasi Error") //akan error karena diberikan panic sama seperti throw new Error javascript
	}

	fmt.Println("Aplikasi berjalan")
}

func endApp(){
	fmt.Println("selesai endApp")
	message := recover();//simpan recover di defer function jangan di function yang terdapat panic
	if(message != nil){
		fmt.Println("Aplikasi berjalan 2", message); //pasang di defer atau function yang dijalankan terakhir
	}
}


//recover : supaya panic tidak menghentikan aplikasinya, atau akan di catch
func runApp2(error bool){
	defer endApp();
	if error {
		panic("Aplikasi Error 2") //akan error karena diberikan panic sama seperti throw new Error javascript
	}

	//jangan pasang panic disini
	// message := recover()
	// fmt.Println("Aplikasi berjalan 2", message)


}

func main(){
	// runApplication(-1)
	runApp2(false);
}


