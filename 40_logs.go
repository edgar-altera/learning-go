package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	// "fmt"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	// log.SetLevel(log.DebugLevel)

    // log.Debug("Useful debugging information.")
    // log.Info("Something noteworthy happened!")
    // log.Warn("You should probably take a look at this.")
    // log.Error("Something failed but I'm not quitting.")

	log.WithFields(
        log.Fields{
            "foo": "foo",
            "bar": "bar",
        },
    ).Info("Something happened")

	p := person { "Edgar", 35, true } 

	log.WithFields((log.Fields{ "name": p.name, "age": p.age, "status": p.status })).Info("Something happened")
}

func init() {

	f, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(f)
}

type person struct {
	name string
	age int
	status bool
}