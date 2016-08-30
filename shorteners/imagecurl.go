package ungo

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func Imagecurl(url_ string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "imgcurl.com"

	var cookies []*http.Cookie

	ageVerificationData := &http.Cookie{ //imagep2p
		Name:   "AgeVerification",
		Path:   "/",
		Domain: "imgcurl.com",
		Value:  "1",
	}

	cookieURL, _ := url.Parse(url_)

	cookies = append(cookies, ageVerificationData)

	cookie.SetCookies(cookieURL, cookies)

	html := htmlDownload(url_, cookie)

	urlregex := regexp.MustCompile(`src="(images/.*?)"`)
	resutl := urlregex.FindAllStringSubmatch(html.Html, -1)[0:]


	return "http://imgcurl.com/" + resutl[0][1], nil
}
