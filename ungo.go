package ungo

import(
	"ungo/shorteners"
	"errors"
)

type Config struct{
	Url string
	Shortener string
}
type Decoded struct{
	UrlDecoded string
	UrlError error
}

var DecodedReturn = new(Decoded)

func Url(info Config)(string,error){

	if info.Shortener == "adfly"{

		DecodedReturn.UrlDecoded,DecodedReturn.UrlError = ungo.Adfly(info.Url)
	}else if info.Shortener == "adfocus"{

		DecodedReturn.UrlDecoded,DecodedReturn.UrlError = ungo.Adfocus(info.Url)
	}else if info.Shortener == "linkbucks"{

		DecodedReturn.UrlDecoded,DecodedReturn.UrlError = ungo.Linkbucks(info.Url)
	}else if info.Shortener == "googl"{

		DecodedReturn.UrlDecoded,DecodedReturn.UrlError = ungo.Googl(info.Url)
	}else{
		DecodedReturn.UrlDecoded,DecodedReturn.UrlError = "",errors.New("The Shortener state is not valid!")
	}

	return DecodedReturn.UrlDecoded,DecodedReturn.UrlError
}