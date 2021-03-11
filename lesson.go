package main

import (
	"fmt"
)


func thirdPartyConnectDB(){
	panic("unable to connect DB!!")
}

func save(){
	defer func(){
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main(){
	save()
	fmt.Println("OK?")

}
