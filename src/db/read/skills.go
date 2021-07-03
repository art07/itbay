package read

import (
	"database/sql"
)

//goland:noinspection GoUnhandledErrorResult
func GetAllSkills() (skills []Skill) {
	itbayDb, _ := sql.Open("sqlite3", "itbay.db")
	defer itbayDb.Close()

	dbRows, _ := itbayDb.Query(`SELECT * FROM skills;`)
	defer dbRows.Close()

	for dbRows.Next() {
		var s Skill
		dbRows.Scan(&s.Id, &s.Skill)
		skills = append(skills, s)
	}

	return
}

//goland:noinspection GoUnhandledErrorResult
func GetSkill(id int) Skill {
	itbayDb, _ := sql.Open("sqlite3", "itbay.db")
	defer itbayDb.Close()

	dbRows, _ := itbayDb.Query(
		`SELECT s.id, s.skill_name, m.id, m.username, m.email 
				FROM skills AS s
				LEFT JOIN members_skills AS ms ON s.id = ms.skill_id
				LEFT JOIN members AS m ON m.id = ms.member_id
				WHERE s.id = ?;`, id)
	defer dbRows.Close()

	var skill Skill
	skill.Members = make([]Member, 0, 8)
	for dbRows.Next() {
		var s Skill
		var m Member
		dbRows.Scan(&s.Id, &s.Skill, &m.Id, &m.UserName, &m.Email)
		if skill.Skill == "" {
			skill = s
		}
		skill.Members = append(skill.Members, m)
	}
	return skill
}
