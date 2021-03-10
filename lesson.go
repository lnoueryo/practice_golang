package main

import (
	"fmt"
)

func foo(params ...int){
	fmt.Println(len(params), params)
}

func main(){
	foo(10, 20)//[10 20]
	foo(10, 20, 30)//[10 20 30]

	s := []int{1, 2, 3}
	foo(s...)//[1 2 3]
}
