# Ungo
Ungo and a package that has the function of unshorten urls.
# How to install
You should run the following commands:<br>
<code>go get github.com/ReiGelado/Ungo</code>
and
<code>go install github.com/ReiGelado/Ungo</code>
<br>There, you can now use :)
# Example
A simple example of how to use the library :)

```go
package main 

import (
	"github.com/ReiGelado/Ungo"
	"fmt"
)

func main() {
	url , err := ungo.Url(ungo.Config{Url:"http://adf.ly/1cXbxn",Shorter:"adfly"})
	if err != nil{
		panic(err)
	}
	fmt.Println(url)
}
```

