package ungo

import(
	"regexp"
)


func Adfocus(url string)(string,error){
	html := htmlDownload(url,"www.adfoc.us")

	urlregex := regexp.MustCompile(`<a href=\"(.*)" class="skip"`)
	resutl := urlregex.FindAllStringSubmatch(html,-1)[0:]

	return resutl[0][1],nil
}