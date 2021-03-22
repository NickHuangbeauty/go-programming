package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello, playground")
	fmt.Println(n)
	fmt.Println(err)

	submain()
}

func submain() {
	n, _ := fmt.Println("This is a test") // ignore err parameter
	fmt.Println(n)
}
