package ungo

import(
	"net/http/cookiejar"
)


func Googl(url string)(string,error){
	cookie , _ := cookiejar.New(nil)

	html := htmlDownload(url,"www.google.com",cookie)

	return html.URL,nil
}