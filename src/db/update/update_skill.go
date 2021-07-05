package update

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//goland:noinspection GoUnhandledErrorResult
func UpdSkill(skillId, newSkillName string) error {
	itbay, err := sql.Open("sqlite3", "itbay.db")
	if err != nil {
		return err
	}
	defer itbay.Close()

	stmt, err := itbay.Prepare(`UPDATE skills SET skill_name = ? WHERE id = ?;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newSkillName, skillId)
	if err != nil {
		return err
	}
	return nil
}
