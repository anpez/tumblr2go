package tumblr2go

import (
	"fmt"
)

// This method returns general information about the blog, such as the title, number of posts, and other high-level data.
func (client *Client) BlogInfo(blogurl string) (*BlogInfo, error) {
	url := fmt.Sprintf("%s/blog/%s/info?api_key=%s", API_BASE_URL, blogurl, client.apiKey)
	var resp struct {
		Blog BlogInfo `json:"blog"`
	}

	err := client.httpClient.Get(url, &resp)
	if nil != err {
		return nil, err
	}

	return &resp.Blog, nil
}

// Retrieve a Blog Avatar
func (client *Client) BlogAvatar(blogurl string, size uint) (string, error) {
	url := fmt.Sprintf("%s/blog/%s/avatar", API_BASE_URL, blogurl)
	if size > 0 {
		url = fmt.Sprintf("%s/%d", url, size)
	}

	var resp struct {
		AvatarURL string `json:"avatar_url"`
	}

	err := client.httpClient.Get(url, &resp)
	if nil != err {
		return "", err
	}

	return resp.AvatarURL, nil
}
