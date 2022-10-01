package http_handler

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
)

func RequestHTTP(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, err
	}
	return res.Body, nil
}
func RequestGZip(url, method string, body io.Reader) ([]byte, error) {
	client := new(http.Client)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept-Encoding", "gzip")

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var f []byte
	if _, err := res.Body.Read(f); err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	g, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return g, nil
}
