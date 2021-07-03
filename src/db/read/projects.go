package read

import (
	"database/sql"
	"log"
)

//goland:noinspection GoUnhandledErrorResult
func GetAllProjects() []Project {
	itbayDb, _ := sql.Open("sqlite3", "itbay.db")
	defer itbayDb.Close()

	dbRows, _ := itbayDb.Query(
		`SELECT p.id AS p_id, p.project_name, p.budget, p.end_till, m.id AS m_id, m.username, m.email
				FROM projects AS p
				LEFT JOIN members AS m ON p.id = m.project_id;`)
	defer dbRows.Close()

	mymap := make(map[string]*[]Member)
	var projects []Project
	for dbRows.Next() {
		var p Project
		var m Member
		dbRows.Scan(&p.Id, &p.ProjectName, &p.BudgetSql, &p.EndTillSql, &m.IdSql, &m.UserNameSql, &m.Email)
		if p.BudgetSql.Valid {
			p.Budget = p.BudgetSql.Int32
		}
		if p.EndTillSql.Valid {
			p.EndTill = p.EndTillSql.String
		}
		if m.IdSql.Valid {
			m.Id = m.IdSql.Int32
		}
		if m.UserNameSql.Valid {
			m.UserName = m.UserNameSql.String
		}

		if _, ok := mymap[p.ProjectName]; !ok {
			members := make([]Member, 0, 8)
			p.Members = &members
			// key = p.ProjectName; value = *[]Member
			mymap[p.ProjectName] = p.Members
			// В *[]Member добавляю member
			if m.UserName != "" {
				*mymap[p.ProjectName] = append(*mymap[p.ProjectName], m)
			}
			projects = append(projects, p)
		} else {
			*mymap[p.ProjectName] = append(*mymap[p.ProjectName], m)
		}

	}

	return projects
}

//goland:noinspection GoUnhandledErrorResult
func GetProject(id int) Project {
	itbayDb, _ := sql.Open("sqlite3", "itbay.db")
	defer itbayDb.Close()

	dbRows, _ := itbayDb.Query(
		`SELECT p.id AS p_id, p.project_name, p.budget, p.end_till, m.id AS m_id, m.username, m.email, md.address, md.itexp_years
				FROM projects AS p
				LEFT JOIN members AS m ON p.id = m.project_id
				LEFT JOIN members_data AS md ON m.id = md.member_id
				WHERE p.id = ?;`, id)
	defer dbRows.Close()

	var pjt Project
	for dbRows.Next() {
		var p Project
		var m Member
		dbRows.Scan(&p.Id, &p.ProjectName, &p.BudgetSql, &p.EndTillSql, &m.IdSql, &m.UserNameSql, &m.Email, &m.AddressSql, &m.ItExpYearsSql)
		if p.BudgetSql.Valid {
			p.Budget = p.BudgetSql.Int32
		}
		if p.EndTillSql.Valid {
			p.EndTill = p.EndTillSql.String
		}
		if m.IdSql.Valid {
			m.Id = m.IdSql.Int32
		}
		if m.UserNameSql.Valid {
			m.UserName = m.UserNameSql.String
		}
		if m.AddressSql.Valid {
			m.Address = m.AddressSql.String
		}
		if m.ItExpYearsSql.Valid {
			m.ItExpYears = m.ItExpYearsSql.Int32
		}

		if pjt.ProjectName == "" {
			members := make([]Member, 0, 8)
			p.Members = &members
			if m.UserName != "" {
				*p.Members = append(*p.Members, m)
			}
			pjt = p
		} else {
			if m.UserName != "" {
				*pjt.Members = append(*pjt.Members, m)
			}
		}
	}
	log.Printf("%#v\n", pjt.GetMembers())
	return pjt
}
