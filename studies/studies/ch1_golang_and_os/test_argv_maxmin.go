package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) == 1{
		fmt.Println("Please give input paramerters.")
	}

	arg := os.Args
	min, err := strconv.ParseFloat(arg[1], 64)
	max := min
	

	for i:=2;i<len(arg);i++{
		n, _ := strconv.ParseFloat(arg[i],64)
		if n < min{
			min = n
		}
		if n> max{
			max = n
		}
	}
	fmt.Println("Min:",min)
	fmt.Println("Max:",max)
}