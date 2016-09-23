package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Smllio(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "smll.io"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`location="(.*)";}`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
