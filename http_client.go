package tumblr2go

import (
	"encoding/json"
	"net/http"
)

type httpClient struct {
	http.RoundTripper
}

func newHttpClient() *httpClient {
	return &httpClient{
		http.DefaultTransport,
	}
}

func (client *httpClient) Get(url string, out interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return err
	}

	resp, err := client.RoundTrip(req)
	if nil != err {
		return err
	}

	defer resp.Body.Close()

	var r Response
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&r)
	if nil != err {
		return err
	}

	return json.Unmarshal(r.Body, out)
}
