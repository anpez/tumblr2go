package tumblr2go

import (
	"bytes"
	"github.com/anpez/tumblr2go/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpClientGet(t *testing.T) {
	const FAKE_URL = "http://theurl.com"

	req, err := http.NewRequest("GET", FAKE_URL, nil)
	assert.Nil(t, err)

	body := bytes.NewBufferString(`{"meta":{"status":200,"msg":"OK"},"response":{"test":"value"}}`)
	resp := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(body),
	}

	mockTransport := new(mocks.RoundTripper)
	mockTransport.On("RoundTrip", req).Return(resp, nil)
	client := httpClient{
		mockTransport,
	}
	assert.NotNil(t, client)

	var out interface{}
	err = client.Get(FAKE_URL, &out)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"test": "value"}, out)
}
