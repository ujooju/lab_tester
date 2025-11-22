package models

type TestRecord struct {
	ID       int    `json:"id"`
	Owner    string `json:"owner"`
	RepoName string `json:"name"`
	Branch   string `json:"branch"`
	Status   string `json:"status"`
	Report   string `json:"report"`
}
