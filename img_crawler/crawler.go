package crawler

import (
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func Images() ([]string, error) {
	var (
		err        error
		imgs       []string
		matches    [][]string
		body       []byte
		content    string
		findImages = regexp.MustCompile("<img.*?src=\"(.*?)\"")
	)
	r, err := http.Get("https://www.ferrari.com/zh-CN/auto/car-range")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return imgs, err
	}
	content = html.UnescapeString(string(body))
	matches = findImages.FindAllStringSubmatch(content, -1)
	for _, val := range matches {
		var imgUrl *url.URL

		// Parse the image URL
		if imgUrl, err = url.Parse(val[1]); err != nil {
			return imgs, err
		}
		if imgUrl.IsAbs() {
			imgs = append(imgs, imgUrl.String())
		}
	}

	return imgs, err
}
