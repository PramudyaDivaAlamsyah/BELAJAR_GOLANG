package main

import "fmt"

type HasName interface{//ini kontrak
	GetName() string//kontrak dari interface
}

func SayHello(hasname HasName){//hanya bisa memasukan parameter yang sudah kontrak dengan HasName
	fmt.Println("hello", hasname.GetName())
}

type Person struct {
	Name string;
}

func (person Person) GetName() string{//lakukan kontrak disini
	return person.Name;
}

//harus membuat implementasi setelah membuat iterface agar berjalan
func main(){
	var dipa Person
	dipa.Name = "DD"
	SayHello(dipa)
}