package tumblr2go

import (
	"encoding/json"
)

type BlogInfo struct {
	Title                string `json:"title"`
	Posts                int    `json:"posts"`
	Name                 string `json:"name"`
	URL                  string `json:"url"`
	Updated              int    `json:"updated"`
	Description          string `json:"description"`
	Ask                  bool   `json:"ask"`
	AskAnon              bool   `json:"ask_anon"`
	Likes                int    `json:"likes"`
	IsBlockedFromPrimary bool   `json:"is_blocked_from_primary"`
}

type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type Response struct {
	Meta Meta            `json:"meta"`
	Body json.RawMessage `json:"response"`
}
