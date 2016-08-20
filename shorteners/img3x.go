package ungo

import (
	"errors"
	"net/http/cookiejar"
	"regexp"
)

func Img3x(url_ string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "img3x.com"

	html := htmlDownload(url_, cookie)

	urlregex := regexp.MustCompile(`<img src="(images/.*)" alt="`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	if resutl == nil {
		return "", errors.New("Image not Found!")
	}

	return "http://img3x.com/" + resutl[0][1], nil
}
