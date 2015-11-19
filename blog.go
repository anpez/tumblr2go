package tumblr2go

import (
	"fmt"
)

// This method returns general information about the blog, such as the title, number of posts, and other high-level data.
func (client *Client) BlogInfo(blogurl string) (*BlogInfo, error) {
	url := fmt.Sprintf("%s/blog/%s/info?api_key=%s", API_BASE_URL, blogurl, client.apiKey)
	var resp struct {
		Meta     meta `json:"meta"`
		Response struct {
			Blog BlogInfo `json:"blog"`
		} `json:"response"`
	}

	err := client.httpClient.Get(url, &resp)
	if nil != err {
		return nil, err
	}

	return &resp.Response.Blog, nil
}
