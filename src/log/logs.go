package main

import (
	"awesomeProject/src/error"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := error.Hello("tim")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
