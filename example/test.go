package main 

import (
	"ungo"
	"fmt"
)

func main() {
	
	url , err := ungo.Shorten(ungo.Config{
		Url:"https://smll.io/9NSF7",
		Shortener:"smllio",
	})

	if err != nil{
		panic(err)
	}

	fmt.Println(url) //GTA-SA TORRENT
} 