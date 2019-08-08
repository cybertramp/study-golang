package main

import(
	"io"
	"os"
)

func main(){
	myString := ""
	arg := os.Args
	if len(arg) == 1{
		myString = "Please give me one argument!"
	}else{
		myString = arg[1]
	}
	io.WriteString(os.Stdout, "This is Standard output!\n")
	io.WriteString(os.Stderr,myString)
	io.WriteString(os.Stderr,"\n")

}