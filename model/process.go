package model

import "amino-server/enums"

type Process struct {
	id                     string
	workflow_id            string
	currentExecutingTaskId string
	processState           enums.ProcessState
}
