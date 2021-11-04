package main

import (
	"fmt"
)

func main() {
	
	es := make(map[string] string) 

	es["StatusUnauthorizedMessage"] = "No Autorizado"

	fmt.Println(es["StatusUnauthorizedMessage"])
}