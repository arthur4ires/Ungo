package ungo

import (
	"net/http/cookiejar"
	"regexp"
)

func Shtio(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "sht.io"
	html := htmlDownload(url, cookie)

	base64 := regexp.MustCompile(`[0-9]/(.*)`)
	resutl := base64.FindAllStringSubmatch(html.URL, -1)[0:]

	b64, err := Base64Decode(resutl[0][1])

	if err != nil {
		panic(err)
	}

	urlregex := regexp.MustCompile(`{sht-io}(.*?){sht-io}`)
	resutl = urlregex.FindAllStringSubmatch(b64,-1)[0:]

	return	resutl[0][1], nil
}
