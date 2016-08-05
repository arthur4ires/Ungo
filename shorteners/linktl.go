package ungo

import (
	"net/http/cookiejar"
	"regexp"
	"strings"
)

const (
	goUrl = "fly/go.php?to="
)

func Linktl(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "link.tl"
	html := htmlDownload(url, cookie)

	token := regexp.MustCompile(".tl/(.*)")
	resutlToken := token.FindAllStringSubmatch(url, -1)[0:]

	HH.Referer = strings.Replace(url, "http://", "", -1)
	htmlGot := htmlDownload(strings.Replace(url, resutlToken[0][1], "", -1)+goUrl+resutlToken[0][1], html.Jar)

	urlregex := regexp.MustCompile(`<div class="skip_btn2"><a href="(.*)">Skip!</a></div>`)
	resutl := urlregex.FindAllStringSubmatch(htmlGot.Html, -1)[0:]

	return resutl[0][1], nil
}
