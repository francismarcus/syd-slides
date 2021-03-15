package main

import "fmt"

type user struct {
	email string
}

func (u user) print() {
	fmt.Println(u.email)
}

func main() {
	u := user{
		email: "test@cygni.se",
	}

	u.print()
}
