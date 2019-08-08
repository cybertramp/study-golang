package main

import(
	"fmt"
	"unsafe"

)

func main(){
	var val int64 = 5
	var p1 = &val	// var의 주소 값
	//var p2 = &p1
	var p2 = (*int32)(unsafe.Pointer(p1))	// p1의 주소를 unsafe

	fmt.Println("*p1:",*p1)		// p1 값의 값 => val의 값
	fmt.Println("*p2:",*p2)		// p2 값의 값 => val의 값(int32로 형변환 됨)

	*p1 = 5434123412312431212	// *p1 값 변경 => val의 값 변경

	fmt.Println("val:",val)
	fmt.Println("*p2:",*p2)		// *p2의 값은 달라졌음(int32로 지정해서 달라짐)
								// 사실 값이 달라진게 아닌 표시를 못하는 거임
	*p1 = 54341234

	fmt.Println("val:",val)
	fmt.Println("*p2:",*p2)		// int32 범위로 하니 정상적으로 돌아옴
}