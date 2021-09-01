package main

import "fmt"

func main() {

	// 1
	switch day:= 1; day {

		case 1:
			fmt.Println("Lunes")
		
		case 2:
			fmt.Println("Martes")
		
		case 3:
			fmt.Println("Miercoles")
		
		case 4:
			fmt.Println("Jueves")
		
		case 5:
			fmt.Println("Viernes")

		case 6:
			fmt.Println("Sábado")

		case 7:
			fmt.Println("Domingo")

		default:
			fmt.Println("Invalid")
	}

	// 2

	var lang int = 2

	switch {
		
		case  lang == 1:
			fmt.Println("Hello")

		case lang == 2:
			fmt.Println("Bonjour")

		case lang == 3:
			fmt.Println("Namstay")
		
		default:
			fmt.Println("Invalid")
	}

	// 3
	
	option:= "five"

	switch option {

		case "one":
			fmt.Println("PHP")
		
		case "two", "three":
			fmt.Println("Javascript")

		case "four", "five", "six":
			fmt.Println("Go")

		default:
			fmt.Println("Invalid")
	}

	// 4

	var value interface {}

	switch q:= value.(type) {
		
		case bool:
			fmt.Println("value is of boolean type")

		case float64:
			fmt.Println("value is of float64 type")

		case int:
			fmt.Println("value is of int type")

		default:
			fmt.Printf("value is of type: %T \n", q)
	}

}