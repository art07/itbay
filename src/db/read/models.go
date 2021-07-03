package read

import "database/sql"

type Member struct {
	IdSql         sql.NullInt32
	Id            int32
	UserNameSql   sql.NullString
	UserName      string
	Email         string
	AddressSql    sql.NullString
	Address       string
	ItExpYearsSql sql.NullInt32
	ItExpYears    int32
	RegData       string
	MemberPjt     Project
	AddInfoSql    sql.NullString
	AddInfo       string
	Skills        []Skill
}

type Project struct {
	Id          int
	ProjectName string
	BudgetSql   sql.NullInt32
	Budget      int32
	EndTillSql  sql.NullString
	EndTill     string
	Members     *[]Member
}

func (p Project) GetMembers() []Member {
	return *p.Members
}

type Skill struct {
	Id      int
	Skill   string
	Members []Member
}
