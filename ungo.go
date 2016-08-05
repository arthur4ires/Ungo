package ungo

import (
	"errors"
	"ungo/shorteners"
)

type Config struct {
	Url       string
	Shortener string
}
type Decoded struct {
	UrlDecoded string
	UrlError   error
}

var DecodedReturn = new(Decoded)

func Url(info Config) (string, error) {

	switch info.Shortener {

	case "adfly":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Adfly(info.Url)
	case "adfocus":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Adfocus(info.Url)
	case "linkbucks":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Linkbucks(info.Url)
	case "googl":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Googl(info.Url)
	case "shst":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Shst(info.Url)
	case "ad7biz":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Ad7biz(info.Url)
	case "tco":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Tco(info.Url)
	case "urlgogs":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Urlgogs(info.Url)
	case "linktl":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Linktl(info.Url)
	case "coegin":
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = ungo.Coegin(info.Url)
	default:
		DecodedReturn.UrlDecoded, DecodedReturn.UrlError = "", errors.New("The Shortener state is not valid!")
	}

	return DecodedReturn.UrlDecoded, DecodedReturn.UrlError
}
