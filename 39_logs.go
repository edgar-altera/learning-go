package main

import (
	// "fmt"
	"log"
	"os"
)

func main() {

	// abrir el archivo webserver.log para escritura
    f, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        log.Fatal(err)
    }
    // y cerrar cuando termine la funcion main
    defer f.Close()
    // asociar el manejador del archivo al log
    log.SetOutput(f)

	log.Fatal("Info", "Testing")
}