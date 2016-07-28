package ungo

import(
	"regexp"
	"net/http/cookiejar"
)


func Adfocus(url string)(string,error){
	cookie , _ := cookiejar.New(nil)

	html := htmlDownload(url,"www.adfoc.us",cookie)

	urlregex := regexp.MustCompile(`<a href=\"(.*)" class="skip"`)
	resutl := urlregex.FindAllStringSubmatch(html.Html,-1)[0:]

	return resutl[0][1],nil
}