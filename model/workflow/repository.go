package workflow

import (
	"amino-server/config"
	"errors"
)

func GetWorkflowByName(name string) (*Workflow, error) {
	db := config.GetDB()
	var workflow Workflow
	if result := db.Where("name=?", "aopusnyhp9amso").First(&workflow); result.Error != nil {
		return nil, errors.New("No Result Found")
	}
	return &workflow, nil
}
