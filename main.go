package main

import (
	"fmt"
	"github.com/MarinX/keylogger"
	"log"
)

func main() {
	devs, err := keylogger.NewDevices()
	if err != nil {
		log.Fatal("oops")
	}
	fmt.Println(devs)
}
