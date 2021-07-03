package init

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//goland:noinspection GoUnhandledErrorResult
func CreateTables() {
	itbayDb, err := sql.Open("sqlite3", "itbay.db")
	if err != nil {
		log.Fatal(err)
	}
	defer itbayDb.Close()

	// 1) Создать таблицу projects.
	stmt, _ := itbayDb.Prepare(`CREATE TABLE IF NOT EXISTS projects (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		project_name TEXT NOT NULL UNIQUE,
		budget INTEGER NULL,
		end_till TEXT NULL
    	);`)
	executeStmt(stmt, "projects")

	// 2) Создать таблицу skills.
	stmt, _ = itbayDb.Prepare(`CREATE TABLE IF NOT EXISTS skills (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		skill_name TEXT NOT NULL UNIQUE
    	);`)
	executeStmt(stmt, "skills")

	// 3) Создать таблицу members.
	stmt, _ = itbayDb.Prepare(`CREATE TABLE IF NOT EXISTS members (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		member_data_id INTEGER NULL UNIQUE,
		project_id INTEGER NULL DEFAULT 1,
        FOREIGN KEY (member_data_id) REFERENCES members_data (id) ON DELETE CASCADE ON UPDATE NO ACTION,
        FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE SET DEFAULT ON UPDATE NO ACTION
    	);`)
	executeStmt(stmt, "members")

	// 4) Создать таблицу members_data.
	stmt, _ = itbayDb.Prepare(`CREATE TABLE IF NOT EXISTS members_data (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		address TEXT NULL,
		itexp_years INTEGER NULL,
		reg_data TEXT NOT NULL,
		member_id INTEGER NOT NULL UNIQUE,
        FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE CASCADE ON UPDATE NO ACTION
    	);`)
	executeStmt(stmt, "members_data")

	// 5) Создать таблицу additional_info.
	stmt, _ = itbayDb.Prepare(`CREATE TABLE IF NOT EXISTS additional_info (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		additional_info TEXT NOT NULL,
		member_id INTEGER NULL UNIQUE,
        FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE CASCADE ON UPDATE NO ACTION
    	);`)
	executeStmt(stmt, "additional_info")

	// 6) Создать таблицу members_skills.
	stmt, _ = itbayDb.Prepare(`CREATE TABLE IF NOT EXISTS members_skills (
		member_id INTEGER NOT NULL,
		skill_id INTEGER NOT NULL,
        FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE CASCADE ON UPDATE NO ACTION,
        FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE ON UPDATE NO ACTION
    	);`)
	executeStmt(stmt, "members_skills")
}

func executeStmt(stmt *sql.Stmt, tableName string) {
	_, err := stmt.Exec()
	if err != nil {
		log.Printf("<%s> table NOT ok", tableName)
	} else {
		log.Printf("<%s> table OK!", tableName)
	}
}
