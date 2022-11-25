package models

type Basic struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

type Columns struct {
	Column  string `json:"column"`
	Type    string `json:"type"`
	Allow   string `json:"allow"`
	Comment string `json:"comment"`
}

type Table struct {
	Name    string    `json:"name"`
	Desc    string    `json:"description"`
	Columns []Columns `json:"columns"`
}

type Describe struct {
	Database Basic   `json:"database"`
	Schema   Basic   `json:"schema"`
	Tables   []Table `json:"tables"`
}
