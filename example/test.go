package main 

import (
	"github.com/ReiGelado/Ungo"
	"fmt"
)

func main() {
	
	url , err := ungo.Shorten(ungo.Config{Url:"http://adf.ly/tYjLr",Shortener:"adfly"})

	if err != nil{
		panic(err)
	}

	fmt.Println(url) //GTA-SA TORRENT
} 