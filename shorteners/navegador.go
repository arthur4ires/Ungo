package ungo

import(
	"net/http"
    "net/http/cookiejar"
    "time"
	"io/ioutil"
)

type HtmlResponse struct{
    Html string
    Jar *cookiejar.Jar
    URL string
}

type HttpHeader struct{
    Host string
    ContentType string
    UserAgent string
    Accept string
    AcceptLanguage string
    CacheControl string
    Referer string
    XRequestedWith string
    Origin string
}

var HH = HttpHeader{
    Host:"www.google.com",
    UserAgent:"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
    Accept:"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    AcceptLanguage:"de,en;q=0.7,en-us;q=0.3",
    CacheControl:"no",
    Referer:"www.google.com",
    ContentType:"application/x-www-form-urlencoded",
    XRequestedWith:"",
    Origin:"",
}

//httpHeader.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko"

func htmlDownload(url string,jar *cookiejar.Jar)HtmlResponse{//(string ,*cookiejar.Jar , string) {
    client := &http.Client{
        Jar: jar,
        Timeout: time.Duration(10 * time.Second),
    }

    req, err := http.NewRequest("GET", url , nil)

    if err != nil {
        panic(err)
    }

    req.Header.Set("Host", HH.Host)
    req.Header.Set("Content-Type",HH.ContentType)
    req.Header.Set("User-Agent", HH.UserAgent)
    req.Header.Set("Accept",HH.Accept)
    req.Header.Set("Accept-Language",HH.AcceptLanguage)
    req.Header.Set("Cache-Control",HH.CacheControl)
    req.Header.Set("Referer",HH.Referer)
    req.Header.Set("Origin",HH.Origin)
    req.Header.Set("X-Requested-With",HH.XRequestedWith)

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
   	}

   	defer resp.Body.Close()
   	body, err := ioutil.ReadAll(resp.Body)
    
    if err != nil {
        panic(err)
    }

    return HtmlResponse{Html:string(body),Jar:jar ,URL : resp.Request.URL.String()}
}