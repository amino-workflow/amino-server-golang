package model

type Workflow struct {
	AuditColumns
	Id      string `json:"id" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Name    string
	Flow    string
	Version int
}
