package ungo

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type HtmlResponse struct {
	Html string
	Jar  *cookiejar.Jar
	URL  string
}

type HttpHeader struct {
	Host           string
	ContentType    string
	UserAgent      string
	Accept         string
	AcceptLanguage string
	CacheControl   string
	Referer        string
	XRequestedWith string
	Origin         string
}

type ResponseJson struct {
	DestinationUrl, Status, Url string
	Success, AdBlockSpotted     bool
}

type Post struct {
	Data url.Values
	Post bool
}

var req *http.Request

var err error

var Rj ResponseJson

var P Post

var HH = HttpHeader{
	Host:           "www.google.com",
	UserAgent:      "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
	Accept:         "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	AcceptLanguage: "de,en;q=0.7,en-us;q=0.3",
	CacheControl:   "no",
	Referer:        "www.google.com",
	ContentType:    "application/x-www-form-urlencoded",
	XRequestedWith: "",
	Origin:         "",
}

func htmlDownload(url string, jar *cookiejar.Jar) HtmlResponse {
	client := &http.Client{
		Jar:     jar,
		Timeout: time.Duration(10 * time.Second),
	}

	if P.Post == false {
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
	} else {
		req, err = http.NewRequest("POST", url, bytes.NewBufferString(P.Data.Encode()))
		if err != nil {
			panic(err)
		}
	}

	req.Header.Set("Host", HH.Host)
	req.Header.Set("Content-Type", HH.ContentType)
	req.Header.Set("User-Agent", HH.UserAgent)
	req.Header.Set("Accept", HH.Accept)
	req.Header.Set("Accept-Language", HH.AcceptLanguage)
	req.Header.Set("Cache-Control", HH.CacheControl)
	req.Header.Set("Referer", HH.Referer)
	req.Header.Set("Origin", HH.Origin)
	req.Header.Set("X-Requested-With", HH.XRequestedWith)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return HtmlResponse{Html: string(body), Jar: jar, URL: resp.Request.URL.String()}
}

func Base64Decode(b64 string) (string, error) {
	b64 = strings.Replace(b64, "=", "", -1) //  illegal base64 data at input byte 56

	data, err := base64.StdEncoding.DecodeString(b64)

	if err != nil {
		panic(err)
	}

	return string(data), nil
}
