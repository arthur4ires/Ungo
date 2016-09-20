package main 

import (
	"ungo"
	"fmt"
)

func main() {
	
	url , err := ungo.Shorten(ungo.Config{
		Url:"http://imgfiles.org/hc5rv10qn9sc/CUTE.gif.html",
		Shortener:"imgfiles",
	})

	if err != nil{
		panic(err)
	}

	fmt.Println(url) //GTA-SA TORRENT
} 