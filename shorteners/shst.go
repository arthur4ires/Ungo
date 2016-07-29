package ungo

import (
	"encoding/json"
	"io"
	"net/http/cookiejar"
	"regexp"
	"strings"
	"time"
)

type ResponseJson struct {
	DestinationUrl, Status string
}

var rj ResponseJson

func Shst(url string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "sh.st"
	HH.Origin = "http://sh.st"
	HH.XRequestedWith = "XMLHttpRequest"

	html := htmlDownload(url, cookie)

	sessionId := regexp.MustCompile(`sessionId: "(.*)",`)
	resutl := sessionId.FindAllStringSubmatch(html.Html, -1)[0:]

	HH.Referer = url

	time.Sleep(5 * time.Second)

	htmlSessionId := htmlDownload("http://sh.st/shortest-url/end-adsession?callback=c&adSessionId="+resutl[0][1], html.Jar)

	jsonResult := htmlSessionId.Html[6:]

	jsonResult = strings.Replace(jsonResult, ");", "", -1)

	dec := json.NewDecoder(strings.NewReader(jsonResult))

	for {
		err := dec.Decode(&rj)
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	return rj.DestinationUrl, nil
}
