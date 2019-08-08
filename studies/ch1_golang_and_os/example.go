/*
[조건]
- 커맨드라인 인수로 입력된 숫자 값들을 모두 더하는 프로그램
- 커맨드라인 인수로 입력된 실수 값에 대한 평균을 구하는 프로그램
- "stop" 이라는 단어를 입력할 때까지 계속해서 입력된 정수 값을 읽는 프로그램
*/

package main

import(
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main(){
	
	var sum float64 = 0

	if len(os.Args) == 1{
		fmt.Println("인자 값을 주세요.")
	}
	arg := os.Args

	n, _ := strconv.ParseFloat(arg[1],64)

	for i:=1;i<len(arg);i++{ // arg[1] 이상 부터
		n, _ = strconv.ParseFloat(arg[i],64)
		sum += n
	}
	// sum
	fmt.Println("Total SUM: ",sum)
	
	// avg
	fmt.Println("Average: ",sum/float64(len(arg)-1))
	
	// check "stop" string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		buf := scanner.Text()
		if buf == "stop"{
			os.Exit(1)
		}
		buf2, err := strconv.Atoi(buf)
		if err != nil{
			fmt.Println("[ERR] This string is not int")
			fmt.Print("> ")
		}else{
			fmt.Println(">",buf2)
		}
		
		
	}
}