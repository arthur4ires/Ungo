package main 

import (
	"ungo"
	"fmt"
)

func main() {
	url , err := ungo.Url(ungo.Config{Url:"",Shortener:""})
	if err != nil{
		panic(err)
	}
	fmt.Println(url)
} 