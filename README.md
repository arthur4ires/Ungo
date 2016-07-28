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
<<<<<<< HEAD
<code>adfly ---> http://adf.ly<br>
adfocus ---> http://adfoc.us<br>
googl ---> http://goo.gl<br></code>
=======
<code>
adfly ---> http://adf.ly <p>
adfocus ---> http://adfoc.us <p>
googl ---> http://goo.gl <p>
</code>
>>>>>>> 71ed9ea06dcd1245d1323d40b4f81fa58fd23bc6

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

