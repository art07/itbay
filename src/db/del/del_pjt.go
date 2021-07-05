package del

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//goland:noinspection GoUnhandledErrorResult
func DeleteProject(pjtId string) error {
	itbay, err := sql.Open("sqlite3", "file:itbay.db?_fk=1")
	if err != nil {
		return err
	}
	defer itbay.Close()

	stmt, err := itbay.Prepare(`DELETE FROM projects WHERE id = ?;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(pjtId)
	if err != nil {
		return err
	}
	return nil
}
