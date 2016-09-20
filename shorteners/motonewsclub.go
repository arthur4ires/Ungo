package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Motonewsclub(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "motonews.club"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`='(.*?)';`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
