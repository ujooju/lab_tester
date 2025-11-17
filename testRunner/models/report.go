package models

type Report struct {
	Score int    `json:"score"`
	Text  string `json:"text"`
}

type TestInfo struct {
	RepoOwner    string `json:"repo_owner"`
	RepoName     string `json:"repo_name"`
	CheckerName  string `json:"checker_name"`
	CheckerToken string `json:"checker_token"`
	Branch       string `json:"branch"`
}
