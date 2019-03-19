package github

import "time"

//IssuesURL URL
const IssuesURL = "https://api.github.com/search/issues"

//IssuesSearchResult result of issues
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

//Issue is for issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User is for user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
