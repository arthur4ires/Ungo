package ungo

import (
	"net/http/cookiejar"
	"regexp"
	"strings"
)

func Coegin(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "coeg.in"
	html := htmlDownload(url, cookie)

	urlregex := regexp.MustCompile(`div class="download-link" style="text-align:center;text-decoration:underline;font-size:20px;"><a href="(.*)" rel="nofollow" target="_blank">`)
	result := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]

	HH.Referer = strings.Replace(url, "http://", "", -1)
	urlredirect := htmlDownload(result[0][1], html.Jar)

	return urlredirect.URL, nil
}
