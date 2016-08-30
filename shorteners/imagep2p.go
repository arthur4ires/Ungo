package ungo

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func Imagep2p(url_ string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "imagep2p.com"

	var cookies []*http.Cookie

	ageVerificationData := &http.Cookie{
		Name:   "AgeVerification",
		Path:   "/",
		Domain: "imagep2p.com",
		Value:  "1",
	}

	cookieURL, _ := url.Parse(url_)

	cookies = append(cookies, ageVerificationData)

	cookie.SetCookies(cookieURL, cookies)

	html := htmlDownload(url_, cookie)

	urlregex := regexp.MustCompile(`src="(images/.*?)"`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:
]
	return "http://imagep2p.com/" + resutl[0][1], nil
}
