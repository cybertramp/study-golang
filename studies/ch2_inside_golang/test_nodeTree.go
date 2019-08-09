/*
go run test_nodeTree.go

# 내부구조 확인
go tool compile -W test_nodeTree.go

go tool compile -W test_nodeTree.go | grep before
go tool compile -W test_nodeTree.go | grep after


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there!")
}
