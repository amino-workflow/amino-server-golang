package workflow

import "amino-server/model"

type Workflow struct {
	model.AuditColumns
	Id      string
	Name    string
	Flow    string
	Version int
}
