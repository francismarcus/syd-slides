package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["route"] = 66
	i := m["route"]
	fmt.Println(i)
}
