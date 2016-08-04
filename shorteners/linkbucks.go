package ungo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	baseUrlIntermission = "/intermission/loadTargetUrl?t="
)

func Linkbucks(url string) (string, error) {
	for i := 0; ; i++ {
		cookie, _ := cookiejar.New(nil)

		HH.Referer = "linkbucks.com"
		HH.Host = ""
		html := htmlDownload(url, cookie)

		urlSite := regexp.MustCompile("AdPopUrl: '(.*)'")
		urlSite2 := regexp.MustCompile("(http://.*?)/")
		tokenUser := regexp.MustCompile("Token: '(.*)'")
		adUrl := regexp.MustCompile(" AdUrl: '(.*)'")
		aK := regexp.MustCompile("params\\['Au' \\+ 'thKey'\\] = (.*);")

		resultUrl := urlSite.FindAllStringSubmatch(html.Html, -1)[0:]
		resultUrlS2 := urlSite2.FindAllStringSubmatch(resultUrl[0][1], -1)[0:]
		resultToken := tokenUser.FindAllStringSubmatch(html.Html, -1)[0:]
		resutladUrl := adUrl.FindAllStringSubmatch(html.Html, -1)[0:]
		resultaK := aK.FindAllStringSubmatch(html.Html, -1)[0:]

		akMid, err := strconv.Atoi(resultaK[0][1][2:][:6])

		if err != nil {
			return "", errors.New("Could not parse the linkbucks authentication key.")
		}

		akNew := resultaK[0][1][:2] + strconv.Itoa(akMid+25) + resultaK[0][1][8:]

		HH.Host = strings.Replace(resultUrlS2[0][1], "http://www.", "", -1)

		htmlAdUrl := htmlDownload(resutladUrl[0][1], html.Jar)

		_ = htmlAdUrl

		time.Sleep(5 * time.Second)

		htmlIntermission := htmlDownload(resultUrlS2[0][1]+baseUrlIntermission+resultToken[0][1]+"&aK="+akNew+"&a_b=false", html.Jar)

		dec := json.NewDecoder(strings.NewReader(htmlIntermission.Html))

		for {
			err := dec.Decode(&Rj)
			if err == io.EOF {
				break
			} else if err != nil {
				return "", err
			}
		}

		if Rj.AdBlockSpotted == true {
			continue
		}

		if Rj.Success == false {
			continue
		} else {
			break
		}
	}

	return Rj.Url, nil
}
