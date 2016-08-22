package ungo

import (
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func Ad2links(url_ string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "ad2links.com"

	//<form method='POST' name='skipform'><input type='submit' class='skip' value='Skip Ad.' name='image'></form>";
	data := url.Values{}
	data.Set("image", "Skip Ad.")

	P.Post = true
	P.Data = data

	html := htmlDownload(url_, cookie)

	urlresult := regexp.MustCompile("window\\.location\\.replace\\('(.*)'\\);")
	resutl := urlresult.FindAllStringSubmatch(html.Html,-1)

	return resutl[0][1], nil
}
