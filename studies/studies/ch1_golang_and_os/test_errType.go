package main

import(
	"errors"
	"fmt"
)

func returnErr(a,b int) error{
	if a==b{
		err := errors.New("Error in returnErr() function!")
		return err
	}else{
		return nil
	}
}

func main(){
	// different value
	err := returnErr(1,2)
	if err == nil{
		fmt.Println("returnErr() ended normaly!")
	}else{
		fmt.Println(err)
	}

	// same value
	err = returnErr(10,10)
	if err == nil{
		fmt.Println("returnErr() ended normaly!")
	}else{
		fmt.Println(err)
	}	

	if err.Error() == "Error in returnErr() function!"{
		fmt.Println("!!")
	}
}