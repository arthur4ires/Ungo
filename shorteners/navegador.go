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

func htmlDownload(url string,host string,jar *cookiejar.Jar)HtmlResponse{//(string ,*cookiejar.Jar , string) {
    client := &http.Client{
        Jar: jar,
        Timeout: time.Duration(10 * time.Second),
    }

    req, err := http.NewRequest("GET", url , nil)

    if err != nil {
        panic(err)
    }

    //req.Header.Set("Host", host)
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko")
    req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    req.Header.Set("Accept-Language", "de,en;q=0.7,en-us;q=0.3")
    req.Header.Set("Cache-Control", "no")

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