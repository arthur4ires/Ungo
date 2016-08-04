package ungo

import (
	"errors"
	"net/http/cookiejar"
	"regexp"
)

func Urlgogs(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "urlgo.gs"
	html := htmlDownload(url, cookie)

	url1regex := regexp.MustCompile(`window\.location="(.*)";}count--;}, 1000\);`)
	url2regex := regexp.MustCompile(`<a href="(.*)" class="btn btn-primary btn-block redirect"`)

	resutl := url1regex.FindAllStringSubmatch(html.Html, -1)[0:]
	result2 := url2regex.FindAllStringSubmatch(html.Html, -1)[0:]

	if resutl[0][1] != result2[0][1] {
		return "", errors.New("Could not parse the urls")
	}

	return resutl[0][1], nil
}
