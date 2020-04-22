package main

// "" and '' is different
import "fmt"

func main() {
	// go does not require ";" at the end, main cannot pass arguments
	fmt.Println("Hello, this is my first go program.")
}

// save and in cmd run: go run hello.go, then compile and execute
