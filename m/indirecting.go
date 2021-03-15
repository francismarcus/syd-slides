package main

import "fmt"

var hello string = "Hello"

func main() {
	s := &hello

	fmt.Println(s)
}
