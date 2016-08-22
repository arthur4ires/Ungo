# Ungo
Ungo is a package that has the function of unshorten urls.
# How to install
You should run the following commands:<br>
<code>go get github.com/ReiGelado/Ungo</code>
and
<code>go install github.com/ReiGelado/Ungo</code>
<br>There, you can now use :)
# Shorteners
The available shorteners are:<br>
<code> 1bebiz ---> http://1be.biz <p>
1tinynet ---> http://1tiny.net <p>
ad7biz ---> http://ad7.biz <p>
adfly ---> http://adf.ly <p>
adfocus ---> http://adfoc.us <p>
amankanlink --> http://kombatch.amankan.link <p>
coegin ---> http://coeg.in <p>
comicon ---> http://comicon.com.br <p>
googl ---> http://goo.gl <p>
gtaind ---> http://i.gtaind.com <p>
hrefli ---> http://href.li <p>
img3x ---> http://img3x.com <p>
linkbucks ---> http://linkbucks.com/<p>
linkshrink ---> http://linkshrink.net/ <p>
linktl ---> http://link.tl <p>
playimg ---> http://playimg.com <p>
shst ---> http://sh.st <p>
shortenid ---> http://shorten.id <p>
tco ---> http://t.co <p>
urlgogs ---> http://urlgo.gs <p>
voyeurimage ---> http://voyeurimage.org <p></code>

# Example
A simple example of how to use the library :)

```go
package main

import (
	"fmt"
	"github.com/ReiGelado/Ungo"
)

func main() {
	url, err := ungo.Shorten(ungo.Config{Url: "http://adf.ly/tYjLr", Shortener:"adfly"})
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}

```
# License
<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Licença Creative Commons" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br /><span xmlns:dct="http://purl.org/dc/terms/" property="dct:title">Ungo</span> está licenciado com uma Licença <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Atribuição-CompartilhaIgual 4.0 Internacional</a>.