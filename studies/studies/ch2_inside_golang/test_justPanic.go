/*

	If you want using panic(), then using with recover()
	Not recommended use only panic()

	And panic() is not same log.Panic()
	log.Panic() send to log system
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Not enough argumnets!")
	}
	fmt.Println("Thanks for the arguments!")
}
