package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Comicon(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "comicon.com.br"
	HH.TimeOut = 30
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`<div class="mP">(.*)</div>`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
