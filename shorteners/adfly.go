package ungo

import (
    "regexp"
    "encoding/base64"
)

var side1 string
var side2 string

func Adfly(url string)(string,error){
    html := htmlDownload(url,"www.adf.ly")

    ysmmregex := regexp.MustCompile("var ysmm = '(.*)';")
    result := ysmmregex.FindAllStringSubmatch(html,-1)[0:]

    for i := 0 ; i < (len(result[0][1]));i++{
        if i%2 == 0{
            side1 += string(result[0][1][i])
        }else{
            side2 = string(result[0][1][i]) + side2
        }
    }

    data, err := base64.StdEncoding.DecodeString(side1+side2)
    if err != nil{
        return "The url can not be decoded!",err
    }

    return string(data[2:]),nil
}