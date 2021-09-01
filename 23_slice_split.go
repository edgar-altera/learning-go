package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := [] byte { '*', 'S', 'a', 'n', 't', ' ', ' ', 'i', 'a', 'g', 'o', '*' }

	fmt.Printf("Slice 1: %s \n", bytes.Split(s1, [] byte(" ")))
}

