package main

import "fmt"

const PI = 3.141516

func main() {

	const APP_ID string = "s92YHAS82jau192"

	example := "Generator"

	union := APP_ID + " " + example

	fmt.Printf("Valor de PI %1.2f - Tipo de Dato: %T\n", PI, PI)
	fmt.Printf("Valor de APP_ID %s - Tipo de Dato: %T\n", APP_ID, APP_ID)
	fmt.Printf("Valor de Union %s - Tipo de Dato: %T\n", union, union)
}