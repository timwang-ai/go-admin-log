package main

import (
	"awesomeProject/src/admin/admin-error"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := admin_error.Hello("tim")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
