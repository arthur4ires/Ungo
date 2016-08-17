package ungo

import (
	"regexp"
)

func Gtaind(url string) (string, error) {
	urlregex := regexp.MustCompile(`\?(.*)`)
	resutl := urlregex.FindAllStringSubmatch(url, -1)[0:]

	decoded, err := Base64Decode(resutl[0][1])

	if err != nil {
		panic(err)
	}

	return decoded, nil
}
