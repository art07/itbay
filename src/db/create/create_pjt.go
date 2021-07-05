package create

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//goland:noinspection GoUnhandledErrorResult
func AddProject(pjtName, pjtBgt, pjtTillDate string) error {
	itbay, err := sql.Open("sqlite3", "itbay.db")
	if err != nil {
		return err
	}
	defer itbay.Close()

	stmt, err := itbay.Prepare(`INSERT INTO projects(project_name, budget, end_till) VALUES (?, ?, ?);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(pjtName, pjtBgt, pjtTillDate)
	if err != nil {
		return err
	}
	return nil
}
