package download

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

func NewClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	client := http.DefaultClient
	client.Jar = jar
	return client
}

func Get(client *http.Client, url string) (io.ReadCloser, error) {
	return download(client, "GET", url)
}

func download(client *http.Client, method string, url string) (io.ReadCloser, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res.Body, nil

}
