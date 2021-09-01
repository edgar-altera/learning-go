package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "***Hello World!***"

	s2 := "$$$Hola $@ @$ Mundo!@@@"

	s3 := "Test, Split, Go!"

	fmt.Printf("Trim s1: %s \n", strings.Trim(s1, "*"))

	fmt.Printf("Trim s2: %s \n", strings.Trim(s2, "@$"))

	fmt.Printf("TrimLeft s1: %s \n", strings.TrimLeft(s1, "*"))

	fmt.Printf("TrimRight s1: %s \n", strings.TrimRight(s1, "*"))

	fmt.Println("SplitAfter s3: ", strings.SplitAfter(s3, ","), len(strings.SplitAfter(s3, ",")))

	split := strings.SplitAfterN(s3, ",", 2)

	for k, v := range split {

		fmt.Printf("Key: %d - Value: %v \n", k, v)
	}

	a := "Hola"
	b := "hola"
	
	fmt.Println(a == b)
	fmt.Println(a != b)

}