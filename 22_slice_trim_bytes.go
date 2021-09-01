package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := [] byte { '*', '*', 'G', 'o', '!', '*', '*' }
	
	s2 := [] byte { '*', 'R', 'u', 'n', 'n', 'i', 'n', 'g', '^' }
	
	s3 := [] byte { '*', 'N', 'o', 'w', '*' }

	fmt.Printf("Slice 1: %s \n", bytes.Trim(s1, "*"))

	fmt.Printf("Slice 2: %s \n", bytes.Trim(s2, "*^"))

	fmt.Printf("Slice 3: %s \n", bytes.Trim(s3, "@"))
}