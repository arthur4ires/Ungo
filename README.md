# Ungo
Ungo and a package that has the function of unshorten urls.
# How to use
<code>
package main <br>
<br>
import (<br>
	"github.com/ReiGelado/Ungo"<br>
	"fmt"<br>
)<br>
<br>
func main() {<br>
	url , err := ungo.Url(ungo.Config{Url:"http://adf.ly/1cXbxn",Shorter:"adfly"})<br>
	if err != nil{<br>
		panic(err)<br>
	}<br>
	fmt.Println(url)<br>
}<br>
</code>
