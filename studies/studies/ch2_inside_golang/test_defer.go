package main

import (
	"fmt"
)

func d1() {
	// for 문 먼저 돌고
	for i := 3; i > 0; i-- {
		// 3 2 1 but defer is LIFO
		defer fmt.Print(i, " ")
	}
	// return 할 때 defer가 동작하면서 함수를 호출하면
	// 그 당시의 변수를 LIFO로 가져오기 때문에 3 2 1 => 1 2 3
	// 으로 가져와 Print
}

func d2() {
	for i := 3; i > 0; i-- {
		// for 루프가 끝나고 익명 함수가 실행
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	// for 루프가 끝나면 i는 0이 됨

	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
