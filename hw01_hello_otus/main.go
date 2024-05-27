package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	otus := "Hello, OTUS!"
	fmt.Println(reverse.String(otus))
}
