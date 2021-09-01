package main

import "fmt"

func main() {

	emp1 := employee { "Edgar", "Development", 35, true }

	emp1.show()

	msg := emp1.message("Saludos!")

	fmt.Println(msg)
}

type employee struct {
	name string
	area string
	age int
	status bool
}

func (e employee) show() {

	fmt.Println("Datos del empleado")
	fmt.Printf("Nombre: %s \n", e.name)
	fmt.Printf("Area: %s \n", e.area)
	fmt.Printf("Edad: %d \n", e.age)
	fmt.Printf("Activo: %t \n", e.status)
}

func (e employee) message(msg string) string {

	resp := "Hola " + e.name + ", tienes un mensaje: " + msg

	return resp
}