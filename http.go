package hackernews

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	urlScheme  = "https"
	urlHost    = "hacker-news.firebaseio.com"
	urlSuffix  = ".json"
	apiVersion = "v0"
)

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildRequestURL(paths ...string) string {
	finalPaths := append([]string{apiVersion}, paths...)
	url := url.URL{
		Scheme: urlScheme,
		Host:   urlHost,
		Path:   path.Join(finalPaths...),
	}
	return url.String() + urlSuffix
}
