package main

import (
	"fmt"
	"os"
)

// Greet prints a greeting message to stdout
func Greet(name string) {
	if name == "" {
		name = "World"
	}
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	Greet(name)
}