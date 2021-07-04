package main

import "fmt"

func main() {
	//hello world
	fmt.Println("hello, world")



	//string
	var nama string = "this a string"
	fmt.Println(nama)

	//const
	const jumlah = 10
	fmt.Println(jumlah)


	//var
	var kota = "jakarta"
	fmt.Println(kota)

	//perbandingan
	var a = 11
	var b = 10
	var c = a == b
	fmt.Println(c)


	//operasi matematika
	var i = 10
	i++
	fmt.Println(i)

	var j = 1
	j += 1
	fmt.Println(j)

	var result = 1220 - 133
	fmt.Println(result)



	//type declaration
	type str string
	type numb int

	var name = "code 12"
	var nomor = 12
	fmt.Println(name)
	fmt.Println(nomor)


	//boolean
	var permission = true
	fmt.Println(permission)

	//conversion
	var desa = "tejoasri"
	var d = desa[0]
	var dDesa = string(d)
	fmt.Println(desa)
	fmt.Println(dDesa)

	//type data number

	var number int = 1341414
	fmt.Println(number)
}