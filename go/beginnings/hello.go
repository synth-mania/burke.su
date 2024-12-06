package main

import "fmt"

func SayHello() {
	// Get a greeting message and print it.
	var message string
	var name string
	var err error
	fmt.Print("Input name: ")
	fmt.Scan(&name)
	message, err = Hello(name)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(message)
	}
}

