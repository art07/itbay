package read

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//goland:noinspection GoUnhandledErrorResult
func GetAllMembers() []Member {
	itbayDb, _ := sql.Open("sqlite3", "itbay.db")
	defer itbayDb.Close()

	dbRows, _ := itbayDb.Query(
		`SELECT m.id, m.username, m.email, md.address, md.itexp_years, md.reg_data, ai.additional_info
				FROM members AS m
				INNER JOIN members_data AS md ON md.id = m.member_data_id
				LEFT JOIN additional_info AS ai ON ai.member_id = m.id
				ORDER BY m.username;`)
	defer dbRows.Close()

	var members []Member
	for dbRows.Next() {
		var m Member
		dbRows.Scan(&m.Id, &m.UserName, &m.Email, &m.AddressSql, &m.ItExpYearsSql, &m.RegData, &m.AddInfoSql)
		if m.AddressSql.Valid {
			m.Address = m.AddressSql.String
		}
		if m.ItExpYearsSql.Valid {
			m.ItExpYears = m.ItExpYearsSql.Int32
		}
		if m.AddInfoSql.Valid {
			m.AddInfo = m.AddInfoSql.String
		}
		members = append(members, m)
	}

	return members
}

//goland:noinspection GoUnhandledErrorResult
func GetMember(id int) Member {
	itbayDb, _ := sql.Open("sqlite3", "itbay.db")
	defer itbayDb.Close()

	var m Member
	itbayDb.QueryRow(
		`SELECT m.username, m.email, md.address, md.itexp_years, md.reg_data, p.id, p.project_name, p.budget, p.end_till, ai.additional_info 
				FROM members AS m
				INNER JOIN members_data AS md ON md.id = m.member_data_id
				INNER JOIN projects AS p ON p.id = m.project_id
				LEFT JOIN additional_info AS ai ON ai.member_id = m.id
				WHERE m.id = ?;`, id).Scan(&m.UserName, &m.Email, &m.AddressSql, &m.ItExpYearsSql, &m.RegData,
		&m.MemberPjt.Id, &m.MemberPjt.ProjectName, &m.MemberPjt.Budget, &m.MemberPjt.EndTillSql, &m.AddInfoSql)
	if m.AddressSql.Valid {
		m.Address = m.AddressSql.String
	}
	if m.ItExpYearsSql.Valid {
		m.ItExpYears = m.ItExpYearsSql.Int32
	}
	if m.MemberPjt.BudgetSql.Valid {
		m.MemberPjt.Budget = m.MemberPjt.BudgetSql.Int32
	}
	if m.MemberPjt.EndTillSql.Valid {
		m.MemberPjt.EndTill = m.MemberPjt.EndTillSql.String
	}
	if m.AddInfoSql.Valid {
		m.AddInfo = m.AddInfoSql.String
	}
	m.Skills = getMemberSkills(itbayDb, id)

	return m
}

//goland:noinspection GoUnhandledErrorResult
func getMemberSkills(itbayDb *sql.DB, id int) (skills []Skill) {
	dbRows, _ := itbayDb.Query(
		`SELECT s.id, s.skill_name FROM members AS m
				INNER JOIN members_skills AS ms ON m.id = ms.member_id
				INNER JOIN skills AS s ON ms.skill_id = s.id
				WHERE m.id = ?;`, id)
	defer dbRows.Close()

	for dbRows.Next() {
		var s Skill
		dbRows.Scan(&s.Id, &s.Skill)
		skills = append(skills, s)
	}
	return
}
