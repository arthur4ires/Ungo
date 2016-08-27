package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Picnictrans(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "picnictrans.com"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`window\.location="(.*?)";}`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
