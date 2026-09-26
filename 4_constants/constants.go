package main

import "fmt"

const age = 24

func main() {
	const name string = "arafat"

	fmt.Println(name)
	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}