package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	var otus string = "Hello, OTUS!"
	fmt.Println(reverse.String(otus))
}
