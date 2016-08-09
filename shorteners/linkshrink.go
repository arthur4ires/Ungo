package ungo

import (
	"errors"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

func Linkshrink(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "linkshrink.net"
	html := htmlDownload(url, cookie)

	urlRedirect := regexp.MustCompile(`<a class='bt' href='(.*)' onClick`)
	result := urlRedirect.FindAllStringSubmatch(html.Html, -1)[0:]

	if result == nil {
		if strings.Contains(html.Html, "Link does not exist") {
			return "", errors.New("The link is invalid!")
		} else {
			return "", errors.New("It was not possible, pull the link.")
		}
	}

	htmlRedirect := htmlDownload(result[0][1], html.Jar)

	return htmlRedirect.URL, nil
}
