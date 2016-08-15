//imgContinue=Continue%20to%20image%20...%20
package ungo

import (
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func Voyeurimage(url_ string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "voyeurimage.org"

	data := url.Values{}
	data.Set("imgContinue", "Continue%20to%20image%20...%20")

	P.Post = true
	P.Data = data

	html := htmlDownload(url_, cookie)

	urlregex := regexp.MustCompile(`<img class='centred' src='(.*)' alt='image' />`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	return resutl[0][1], nil
}
