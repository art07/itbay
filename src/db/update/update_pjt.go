package update

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//goland:noinspection GoUnhandledErrorResult
func UpdPjt(pjtId, newPjtName, newPjtBgt, newTillTimeStr string) error {
	itbay, err := sql.Open("sqlite3", "itbay.db")
	if err != nil {
		return err
	}
	defer itbay.Close()

	stmt, err := itbay.Prepare(
		`UPDATE projects 
				SET project_name = ?, 
					budget = ?, 
					end_till = ? 
				WHERE id = ?;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newPjtName, newPjtBgt, newTillTimeStr, pjtId)
	if err != nil {
		return err
	}
	return nil
}
