package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Imagebam(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "imagebam.com"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`<a href="(.*)?download=`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
