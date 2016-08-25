package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Ppw(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "p.pw"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`window\.location = "(.*)"`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
