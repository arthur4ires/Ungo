# Ungo
Ungo and a package that has the function of unshorten urls.
# How to use
<pre><code>
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
</code></pre>
