package init

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//goland:noinspection GoUnhandledErrorResult,SqlInsertValues
func CreateEntries() {
	itbayDb, err := sql.Open("sqlite3", "itbay.db")
	if err != nil {
		log.Fatal(err)
	}
	defer itbayDb.Close()

	// 1) Создать записи для таблицы skills.
	stmt, _ := itbayDb.Prepare("INSERT INTO skills(skill_name) VALUES (?)")
	insertDataToTable(stmt, &skillsTable)

	// 2) Создать записи для таблицы projects.
	stmt, _ = itbayDb.Prepare("INSERT INTO projects(project_name, budget, end_till) VALUES (?, ?, ?)")
	insertDataToProjects(stmt, &pjtTblInfo)

	// 3) Создать записи для таблицы additional_info.
	stmt, _ = itbayDb.Prepare("INSERT INTO additional_info(additional_info) VALUES (?)")
	insertDataToTable(stmt, &additionalInfoTable)

	// 4) Создать записи для таблицы members.
	stmt, _ = itbayDb.Prepare("INSERT INTO members(username) VALUES (?)")
	insertDataToTable(stmt, &membersTable)

	// 5) Создать записи для таблицы members_data.
	// В ручном режиме.

	// 6) Создать записи для таблицы members_skills.
	// В ручном режиме.
}

func insertDataToTable(stmt *sql.Stmt, tblInfo *tableInfo) {
	for _, item := range (*tblInfo).DataArr {
		_, err := stmt.Exec(item)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Record <%s> added to <%s>\n", item, (*tblInfo).TableName)
		}
	}
}

func insertDataToProjects(stmt *sql.Stmt, pjtTblInfo *projectsTableInfo) {
	for _, item := range (*pjtTblInfo).ProjectsArr {
		_, err := stmt.Exec(item.Name, item.Budget, item.EndTill)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Record <%#v> added to <%s>\n", item, (*pjtTblInfo).TableInfo.TableName)
		}
	}
}
