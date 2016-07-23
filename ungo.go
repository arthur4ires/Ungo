package ungo

type Config struct{
	Url string
	Shorter string
}
type Decoded struct{
	UrlDecoded string
	UrlError error
}

var DecodedReturn = new(Decoded)

func Url(info Config)(string,error){
	if info.Shorter == "adfly"{
		DecodedReturn.UrlDecoded,DecodedReturn.UrlError = Adfly(info.Url)

	}
	return DecodedReturn.UrlDecoded,DecodedReturn.UrlError
}