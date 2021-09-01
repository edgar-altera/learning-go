package main

import "fmt"

func main()  {
	
	var number_int int8
	var number_float float32
	var name  string
	var status bool
	var sports [] string
	var langs [3] string
	var x, y, z int = 2, 4, 6
	var slice [] int 

	country:= "Chile"
	number:= 1
	front, back, db := "React", "Go", "MongoDB"
	langs[0] = front
	langs[1] = back
	langs[2] = db

	fmt.Printf("Número entero: %d Tipo de dato: %T\n" , number_int, number_int)
	fmt.Printf("Número decimal: %.2f Tipo de dato: %T\n" , number_float, number_float)
	fmt.Printf("Cadena: %s Tipo de dato: %T\n" , name, name)
	fmt.Printf("Status: %t Tipo de dato: %T\n" , status, status)
	fmt.Printf("Slice: %v Tipo de dato: %T\n" , sports, sports)
	fmt.Printf("Array: %v Tipo de dato: %T\n" , langs, langs)
	fmt.Printf("Numeros: %d %d %d\n", x, y, z)
	fmt.Printf("Pais: %s\n", country)
	fmt.Printf("Num: %d Tipo de dato: %T\n", number, number)
	fmt.Printf("Frontend: %s Backend: %s DB: %s \n", front, back, db)
	fmt.Printf("Slice is nil? %t \n", slice == nil) 

}