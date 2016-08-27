//http://gadget14.pw/compare-ability-xiaomi-mi-5-vs-iphone-6s/?d=aHR0cDovL3d3dzY3LnppcHB5c2hhcmUuY29tL3YvOFJKUm9hR2QvZmlsZS5odG1s#landing
package ungo

import (
	"regexp"
)

func Gadget14(url string) (string, error) {
	urlregex := regexp.MustCompile(`d=(.*)[#]`)
	resutl := urlregex.FindAllStringSubmatch(url, -1)[0:]

	decoded, err := Base64Decode(resutl[0][1])

	if err != nil {
		panic(err)
	}

	return decoded, nil
}
