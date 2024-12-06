package main

import ("fmt"
	"errors")

func Hello(name string) (string, error) {
	// return a greeting that embeds the name in a message

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
