package main

import "fmt"

func main() {
	// Control Flow:
	// 1. Sequence
	// 2. for loop; iterative
	// 3. conditions

	// This is test fmt Println()
	fmt.Println("This is my first time to learn GOlang")

	foo()

	bar()
}

func foo() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func bar() {
	fmt.Println("This is the end of for loop")
}
