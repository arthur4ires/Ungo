package ungo

import (
	"net/http/cookiejar"
)

func Tco(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "t.co"
	html := htmlDownload(url, cookie)

	return html.URL, nil
}
