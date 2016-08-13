package ungo

import (
	"errors"
	"net/http/cookiejar"
	"regexp"
)

var side1 string
var side2 string

func Adfly(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "adf.ly"
	html := htmlDownload(url, cookie)

	if html.URL == "http://adf.ly/not-found.php" {
		return "", errors.New("The URL is not a ADFLY valid link...")
	}

	ysmmregex := regexp.MustCompile("var ysmm = '(.*)';")
	result := ysmmregex.FindAllStringSubmatch(html.Html, -1)[0:]

	if result == nil {
		return "", errors.New(url + " is not a adfly valid link...")
	}

	for i := 0; i < (len(result[0][1])); i++ {
		if i%2 == 0 {
			side1 += string(result[0][1][i])
		} else {
			side2 = string(result[0][1][i]) + side2
		}
	}

	data, err := Base64Decode(side1 + side2)
	if err != nil {
		panic(err)
	}

	return string(data[2:]), nil
}
