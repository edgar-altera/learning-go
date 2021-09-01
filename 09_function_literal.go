package main

import "fmt"

func main() {

	func() { fmt.Println("Welcome to Go!") }()

	greeting := func(name string) { fmt.Printf("Hi %s, Welcome! \n", name) }

	greeting("Edgar")

	passfunc(greeting)

	saludar := returnfunc()

	saludar("Caleb")

}

func passfunc(saludo func(nombre string)) {
		
	saludo("Josué")
}

func returnfunc() func(name string) {

	myf := func(name string) { fmt.Printf("Hola %s, Bienvenido! \n", name)}

	return myf
}


