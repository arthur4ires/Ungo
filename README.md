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
ad2links ---> https://ad2links.com <p>
ad7biz ---> http://ad7.biz <p>
adfly ---> http://adf.ly <p>
adfocus ---> http://adfoc.us <p>
amankanlink --> http://kombatch.amankan.link <p>
coegin ---> http://coeg.in <p>
comicon ---> http://comicon.com.br <p>
freeanimeonline ---> http://freeanimeonline.xyz <p>
gadget14 ---> http://gadget14.pw <p>
googl ---> http://goo.gl <p>
gtaind ---> http://i.gtaind.com <p>
hrefli ---> http://href.li <p>
imagebam ---> http://imagebam.com <p>
img3x ---> http://img3x.com <p>
imagecurl ---> http://imgcurl.com <p>
imagep2p ---> http://imagep2p.com <p>
linkbucks ---> http://linkbucks.com/<p>
linkshrink ---> http://linkshrink.net/ <p>
linktl ---> http://link.tl <p>
motonewsclub ---> http://motonews.club <p>
picnictrans ---> http://www.picnictrans.com/ <p>
playimg ---> http://playimg.com <p>
ppw ---> http://p.pw <p>
shst ---> http://sh.st <p>
shtio ---> http://sht.io <p>
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
Ungo is under the MIT license.
You can read the full terms here: <a href = "https://github.com/ReiGelado/Ungo/blob/master/LICENSE">MIT</a>