package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			// 20, 40, 60, 80
			// Println() will be not work!
			continue
		}

		if i == 95 {
			// 95 will be stopped.
			// So Println() will be not work!
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 10
	for { // 10 9 8 7 6 5 4 3 2 1 0
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	i = 0
	expression := true
	for ok := true; ok; ok = expression {
		if i > 10 {
			expression = false
		}

		fmt.Println(i, " ")
		i++
	}
	fmt.Println()

	arr := [5]int{0, 1, -1, 2, -2}
	for i, val := range arr {
		fmt.Println("index:", i, "value:", val)
	}

	two := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(two)

	three := [2][2][2]int{
		{{1, 0}, {-2, 4}},
		{{5, -1}, {7, 0}},
	}
	fmt.Println(three)
}
