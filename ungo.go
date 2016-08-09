package ungo

import (
	"errors"
	"strings"
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

type function func(string) (string, error)

func tratament(y string) string {
	return strings.Replace(y, " ", "", -1)
}

func Shorten(info Config) (string, error) {
	info.Shortener = tratament(info.Shortener) //spaces

	shortenersMap := map[string]function{
		"adfly":       ungo.Adfly,
		"ad7biz":      ungo.Ad7biz,
		"adfocus":     ungo.Adfocus,
		"amankanlink": ungo.Amankanlink,
		"coegin":      ungo.Coegin,
		"googl":       ungo.Googl,
		"linkbucks":   ungo.Linkbucks,
		"linkshrink":  ungo.Linkshrink,
		"linktl":      ungo.Linktl,
		"shst":        ungo.Shst,
		"tco":         ungo.Tco,
		"urlgogs":     ungo.Urlgogs,
	}

	if val, ok := shortenersMap[info.Shortener]; ok {
		return val(info.Url)
	} else {
		return "", errors.New("The Shortener state is not valid!")
	}
}
