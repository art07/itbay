package create

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//goland:noinspection GoUnhandledErrorResult
func AddSkill(skillName string) error {
	itbay, err := sql.Open("sqlite3", "itbay.db")
	if err != nil {
		return err
	}
	defer itbay.Close()

	stmt, err := itbay.Prepare(`INSERT INTO skills(skill_name) VALUES (?);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(skillName)
	if err != nil {
		return err
	}
	return nil
}
