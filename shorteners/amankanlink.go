package ungo

import (
	"encoding/base64"
	"errors"
	"regexp"
)

func Amankanlink(url string) (string, error) {
	urlregex := regexp.MustCompile(`/p/go\.html\?url=(.*)`)
	resutl := urlregex.FindAllStringSubmatch(url, -1)[0:]

	data, err := base64.StdEncoding.DecodeString(resutl[0][1])
	if err != nil {
		return "", errors.New("The url can not be decoded!")
	}

	return string(data), nil
}
