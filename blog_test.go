package tumblr2go

import (
	"encoding/json"
	"fmt"
	"github.com/anpez/tumblr2go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	FAKE_API_KEY  = "apikey"
	FAKE_BLOG_URL = "myurl.tumblr.com"
)

func TestBlogInfo(t *testing.T) {
	mockHttp := new(mocks.HttpClient)
	mockHttp.On("Get", fmt.Sprintf("https://api.tumblr.com/v2/blog/%s/info?api_key=%s", FAKE_BLOG_URL, FAKE_API_KEY), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		resp := `{"blog":{"title":"David's Log","posts":3456,"name":"david","url":"http://david.tumblr.com/","updated":1308953007,"description":"<p><strong>Mr. Karp</strong> is tall and skinny, with unflinching blue eyes a mop of brown hair. He speaks incredibly fast and in complete paragraphs.</p>","ask":true,"ask_anon":false,"likes":12345}}`

		assert.Nil(t, json.Unmarshal([]byte(resp), args.Get(1)))
	})

	client := NewClientWithTransport(FAKE_API_KEY, mockHttp)
	assert.NotNil(t, client)
	info, err := client.BlogInfo(FAKE_BLOG_URL)

	assert.Nil(t, err)
	assert.NotNil(t, info)

	validInfo := &BlogInfo{
		Title:                "David's Log",
		Posts:                3456,
		Name:                 "david",
		URL:                  "http://david.tumblr.com/",
		Updated:              1308953007,
		Description:          `<p><strong>Mr. Karp</strong> is tall and skinny, with unflinching blue eyes a mop of brown hair. He speaks incredibly fast and in complete paragraphs.</p>`,
		Ask:                  true,
		AskAnon:              false,
		Likes:                12345,
		IsBlockedFromPrimary: false,
	}
	assert.Equal(t, info, validInfo)
}

func TestBlogAvatar(t *testing.T) {
	mockHttp := new(mocks.HttpClient)
	mockHttp.On("Get", fmt.Sprintf("https://api.tumblr.com/v2/blog/%s/avatar", FAKE_BLOG_URL), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		resp := `{"avatar_url":"https:\/\/38.media.tumblr.com\/avatar_fdf0635a9d74_64.png"}`

		assert.Nil(t, json.Unmarshal([]byte(resp), args.Get(1)))
	})
	mockHttp.On("Get", fmt.Sprintf("https://api.tumblr.com/v2/blog/%s/avatar/128", FAKE_BLOG_URL), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		resp := `{"avatar_url":"https:\/\/38.media.tumblr.com\/avatar_fdf0635a9d74_128.png"}`

		assert.Nil(t, json.Unmarshal([]byte(resp), args.Get(1)))
	})

	client := NewClientWithTransport(FAKE_API_KEY, mockHttp)
	assert.NotNil(t, client)

	avatarURL, err := client.BlogAvatar(FAKE_BLOG_URL, 0)
	assert.Nil(t, err)
	assert.NotNil(t, avatarURL)
	assert.Equal(t, avatarURL, "https://38.media.tumblr.com/avatar_fdf0635a9d74_64.png")

	avatarURL, err = client.BlogAvatar(FAKE_BLOG_URL, 128)
	assert.Nil(t, err)
	assert.NotNil(t, avatarURL)
	assert.Equal(t, avatarURL, "https://38.media.tumblr.com/avatar_fdf0635a9d74_128.png")
}
