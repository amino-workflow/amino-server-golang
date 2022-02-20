package migrations

import (
	"amino-server/config"
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
	db := config.GetDB()
	db.Exec(string(query))
}
