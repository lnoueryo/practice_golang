package main

import (
	"fmt"
)

func one(x *int){
	*x = 1
}


func main(){
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(&n)

	// fmt.Println(&n)
	// var f = 100
	// fmt.Println(&f)
	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)
}
