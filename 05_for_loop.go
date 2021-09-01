package main

import "fmt"

func main() {


	// 1

	for i := 0; i < 4; i++ {

		fmt.Printf("Go in Altera - i = %d \n", i)
	}

	// 2 

	// for {
	// 	// fmt.Println("Hello Infinite World! ")
	// }

	// 3 

	i:= 0
	for i < 3 {
		i += 1
		fmt.Printf("Go - i = %d \n", i)
	}

	// 4

	langs := [] string { "PHP", "JavaScript", "Go" }

	for key, value := range langs {
		
		fmt.Printf("Key: %d - Value: %s \n", key, value)
	}

	// 5

	developer := map [string] string {
		"name": "Edgar",
		"frontend": "React",
		"backend": "Go",
	}

	for key, value := range developer {
		fmt.Printf("Key: %s - Value: %s \n", key, value)
	}

	// 6 

	// using channel
    chnl := make(chan int)
    go func(){
        chnl <- 100
        chnl <- 1000
       	chnl <- 10000
       	chnl <- 100000
       	close(chnl)
    }()
	
    for i:= range chnl {
       fmt.Println(i) 
    }

}