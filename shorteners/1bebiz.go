package ungo

import (
	"regexp"
)

func Bebiz(url string) (string, error) {
	urlregex := regexp.MustCompile(`\?(.*)`)
	resutl := urlregex.FindAllStringSubmatch(url, -1)[0:]

	return resutl[0][1], nil
}
