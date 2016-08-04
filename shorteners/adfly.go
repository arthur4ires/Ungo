package ungo

import (
	"encoding/base64"
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

	ysmmregex := regexp.MustCompile("var ysmm = '(.*)';")
	result := ysmmregex.FindAllStringSubmatch(html.Html, -1)[0:]

	if result == nil{
		return "",errors.New(url + " is not a adfly valid link...")
	}

	for i := 0; i < (len(result[0][1])); i++ {
		if i%2 == 0 {
			side1 += string(result[0][1][i])
		} else {
			side2 = string(result[0][1][i]) + side2
		}
	}

	data, err := base64.StdEncoding.DecodeString(side1 + side2)
	if err != nil {
		return "", errors.New("The url can not be decoded!")
	}

	return string(data[2:]), nil
}
