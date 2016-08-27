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
		"1bebiz":	       ungo.Bebiz,
		"1tinynet":        ungo.Tinynet,
		"adfly":       	   ungo.Adfly,
		"ad2links":        ungo.Ad2links,
		"ad7biz":      	   ungo.Ad7biz,
		"adfocus":         ungo.Adfocus,
		"amankanlink":     ungo.Amankanlink,//"bcvc":	 	   ungo.Bcvc, //not implemented
		"coegin":          ungo.Coegin,
		"comicon":         ungo.Comicon,
		"freeanimeonline": ungo.Freeanimeonline,
		"gadget14":		   ungo.Gadget14,	
		"googl":           ungo.Googl,
		"gtaind":          ungo.Gtaind,
		"hrefli":          ungo.Hrefli,
		"img3x":           ungo.Img3x,
		"imagep2p":        ungo.Imagep2p,
		"linkbucks":       ungo.Linkbucks,
		"linkshrink":      ungo.Linkshrink,
		"linktl":          ungo.Linktl,
		"picnictrans":     ungo.Picnictrans,
		"playimg":         ungo.Playimg,
		"ppw":             ungo.Ppw,
		"shst":            ungo.Shst,
		"shtio":	       ungo.Shtio,
		"shortenid":       ungo.Shortenid,
		"tco":             ungo.Tco,
		"urlgogs":         ungo.Urlgogs,
		"voyeurimage":     ungo.Voyeurimage,
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
