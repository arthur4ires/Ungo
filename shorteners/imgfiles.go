package ungo

import (
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func ImgFiles(url_ string) (string, error) {
	cookie, _ := cookiejar.New(nil)

	HH.Host = "imgfiles.org"

	//
	//	<Form name="F1" method="POST" action=''>
	//	<input type="hidden" name="op" value="view">
	//	<input type="hidden" name="id" value="hc5rv10qn9sc">
	//	<input type="hidden" name="pre" value="1">
	//	<!--
	//	<input type="submit" name="next" value="Continue to image...">
	//	-->
	//	</Form>

	html := htmlDownload(url_, cookie)

	idResult := regexp.MustCompile(`name="id" value="(.*)">`)
	result := idResult.FindAllStringSubmatch(html.Html,-1)

	data := url.Values{}
	data.Set("id", result[0][1])
	data.Add("op","view")
	data.Add("pre","1")

	P.Post = true
	P.Data = data

	htmlP := htmlDownload(url_,html.Jar)	

	imageUrl := regexp.MustCompile(`<img src="(.*?)"`)
	resultUrl := imageUrl.FindAllStringSubmatch(htmlP.Html,-1)

	return resultUrl[0][1] , nil
}
