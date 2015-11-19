package tumblr2go

import (
	"encoding/json"
	"net/http"
)

type httpClient struct {
	http.Client
}

func newHttpClient() *httpClient {
	return new(httpClient)
}

func (client *httpClient) Get(url string, out interface{}) error {
	response, err := client.Client.Get(url)
	if nil != err {
		return err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	return decoder.Decode(out)
}
