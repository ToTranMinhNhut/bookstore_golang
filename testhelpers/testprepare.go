package testhelpers

import (
	"bookstoreupdate/db"
	"io/ioutil"
	"strings"
)

// Prepare for test read .sql file and execute it
func PrepareDBForTest(db *db.DB, path string) error {
	// Read .sql file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Split statements in .sql file
	requests := strings.Split(string(file), ";")

	// Execute sql statements
	for _, request := range requests {
		_, err := db.Client.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}
