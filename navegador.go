package ungo

import(
	"net/http"
	"io/ioutil"
)

func htmlDownload(url string,host string) string {
    client := &http.Client{}

    req, err := http.NewRequest("GET", url , nil)
    if err != nil {
        panic(err)
    }

    req.Header.Set("Host", host)
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

    return string(body)
}