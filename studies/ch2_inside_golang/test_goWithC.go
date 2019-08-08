package main

//#include<stdio.h>
//void callC(){
//	printf("Clling C code!\n");
//}

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}
