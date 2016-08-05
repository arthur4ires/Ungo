# Ungo
Ungo and a package that has the function of unshorten urls.
# How to install
You should run the following commands:<br>
<code>go get github.com/ReiGelado/Ungo</code>
and
<code>go install github.com/ReiGelado/Ungo</code>
<br>There, you can now use :)
# Shorteners
The available shorteners are:<br>
<code> ad7biz ---> http://ad7.biz <p>
adfly ---> http://adf.ly <p>
adfocus ---> http://adfoc.us <p>
googl ---> http://goo.gl <p>
linkbucks ---> http://linkbucks.com/<p>
shst ---> http://sh.st <p>
tco ---> http://t.co <p>
urlgogs ---> http://urlgo.gs <p>
linktl ---> http://link.tl <p></code>

# Example
A simple example of how to use the library :)

```go
package main 

import (
	"github.com/ReiGelado/Ungo"
	"fmt"
)

func main() {
	url , err := ungo.Url(ungo.Config{Url:"http://adf.ly/1cXbxn",Shortener:"adfly"})
	if err != nil{
		panic(err)
	}
	fmt.Println(url)
}
```