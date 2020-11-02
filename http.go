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
	apiVersion = "v0"
)

func get(url string) ([]byte, error) {
	// Call endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Extract response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildRequestURL(paths ...string) string {
	finalPaths := []string{apiVersion}
	finalPaths = append(finalPaths, paths...)
	url := url.URL{
		Scheme: urlScheme,
		Host:   urlHost,
		Path:   path.Join(finalPaths...),
	}
	return url.String() + ".json"
}
