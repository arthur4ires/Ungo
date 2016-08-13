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

type function func(string) (string, error)

func Shorten(info Config) (string, error) {
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
		"shortenid":   ungo.Shortenid,
		"tco":         ungo.Tco,
		"urlgogs":     ungo.Urlgogs,
	}

	if info.Shortener == "" {
		return "", errors.New("Shortener can not be empty ...")
	}

	if urlReturn, ok := shortenersMap[info.Shortener]; ok {
		return urlReturn(info.Url)
	} else {
		return "", errors.New("The Shortener state is not valid!")
	}
}
