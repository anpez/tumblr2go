package tumblr2go

import (
	"encoding/json"
	"fmt"
	"github.com/anpez/tumblr2go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBlogInfo(t *testing.T) {
	const API_KEY = "apikey"
	const BLOG_URL = "myurl.tumblr.com"

	mockHttp := new(mocks.HttpClient)
	mockHttp.On("Get", fmt.Sprintf("https://api.tumblr.com/v2/blog/%s/info?api_key=%s", BLOG_URL, API_KEY), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		resp := `{"meta":{"status":200,"msg":"OK"},"response":{"blog":{"title":"David's Log","posts":3456,"name":"david","url":"http:\/\/david.tumblr.com\/","updated":1308953007,"description":"<p><strong>Mr. Karp<\/strong> is tall and skinny, with unflinching blue eyes a mop of brown hair.\r\nHe speaks incredibly fast and in complete paragraphs.</p>","ask":true,"ask_anon":false,"likes":12345}}}`

		assert.Nil(t, json.Unmarshal([]byte(resp), args.Get(1)))
	})

	client := NewClientWithHttp(API_KEY, mockHttp)
	assert.NotNil(t, client)
	info, err := client.BlogInfo(BLOG_URL)

	assert.Nil(t, err)
	assert.NotNil(t, info)
}
