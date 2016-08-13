package ungo

import (
	"regexp"
)

func Shortenid(url string) (string, error) {
	urlregex := regexp.MustCompile(`\?url=(.*)`)
	resutl := urlregex.FindAllStringSubmatch(url, -1)[0:]

	url, err := Base64Decode(resutl[0][1])

	if err != nil {
		panic(err)
	}

	return url, nil
}
