package ungo

import (
    "regexp"
    "encoding/base64"
)

var lado1 string
var lado2 string

func Adfly(url string)(string,error){
    html := htmlDownload(url)

    ysmmregex := regexp.MustCompile("var ysmm = '(.*)';")
    resultado := ysmmregex.FindAllStringSubmatch(html,-1)[0:]

    for i := 0 ; i < (len(resultado[0][1]));i++{
        if i%2 == 0{
            lado1 += string(resultado[0][1][i])
        }else{
            lado2 = string(resultado[0][1][i]) + lado2
        }
    }

    data, err := base64.StdEncoding.DecodeString(lado1+lado2)
    if err != nil{
        return "A url não pode ser decodificada!",err
    }

    return string(data[2:]),nil
}