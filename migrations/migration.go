package migrations

import (
	"amino-server/config"
	"amino-server/model"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// getAbsoluteFile -> gets absolute path of config
func getAbsoluteFile(relativeFilePath string) string {
	// Get absolute file path of the config file
	filePath, err := filepath.Abs(relativeFilePath)
	if err != nil {
		panic("Unable to read config")
	}

	return filePath
}

func Migrate() {
	query, err := ioutil.ReadFile(getAbsoluteFile("migrations/changelogs/changelog-v0.01-create-table.sql"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(query))
	db := config.GetDB()
	db.Exec(string(query))

	var workflow model.Workflow
	db.Where("name=?", "aopusnyhp9aso").Find(&workflow)
	fmt.Println(workflow)
}
