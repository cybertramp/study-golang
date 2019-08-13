package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()") // 1
	defer func() {            // 5 a()가 return 하려는 시점
		if c := recover(); c != nil { // 6
			fmt.Println("Recover inside a()!") // 7
		}
	}()

	fmt.Println("About to call b()") // 2
	b()                              // 3
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()") // 4
	panic("Panic in b()!")    // 5
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!") // 8
}
