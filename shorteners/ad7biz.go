package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Ad7biz(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "ad7.biz"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`&oid=0&url=(.*)&ref=`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
