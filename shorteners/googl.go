package ungo

import (
	"net/http/cookiejar"
)

func Googl(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "goo.gl"
	html := htmlDownload(url, cookie)

	return html.URL, nil
}
