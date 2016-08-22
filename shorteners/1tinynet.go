package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Tinynet(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "tiny.net"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`href="(.*)">click here`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
