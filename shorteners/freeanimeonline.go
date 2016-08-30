//www.freeanimeonline.xyz/top-tips-health-insurance-need-know/?site=aHR0cDovL3d3dy5zb2xpZGZpbGVzLmNvbS9kL2JjNjA0ZmI3MjMv&c=0&user=41748
package ungo

import (
	"regexp"
)

func Freeanimeonline(url string) (string, error) {
	urlregex := regexp.MustCompile(`site=(.*?)&c=`)
	resutl := urlregex.FindAllStringSubmatch(url, -1)[0:]

	decoded, err := Base64Decode(resutl[0][1])

	if err != nil {
		panic(err)
	}

	return decoded, nil
}
