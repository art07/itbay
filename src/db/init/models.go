package init

type tableInfo struct {
	TableName string
	DataArr   []string
}

type project struct {
	Name    string
	Budget  int
	EndTill string
}

type projectsTableInfo struct {
	TableInfo   tableInfo
	ProjectsArr []project
}
