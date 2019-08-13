/* argv float type check program */
package main

import(
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) == 1{
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arg := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64

	// check arg float
	for err != nil{
		if k >= len(arg){
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arg[k], 64)
		k++
	}

	min, max := n,n

	for i:=2;i<len(arg);i++{
		n, err := strconv.ParseFloat(arg[i], 64)
		if err == nil{
			if n < min{
				min = n
			}

			if n > max{
				max = n
			}
		}
		
	}
	fmt.Println("Min:",min)
	fmt.Println("Max:",max)
}