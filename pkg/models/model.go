package models

type Columns struct {
	Table   string `json:"table"`
	Column  string `json:"column"`
	Type    string `json:"type"`
	Allow   string `json:"allow"`
	Comment string `json:"comment"`
}

type Describe struct {
	Database string `json:"database"`
}
