package init

var skillsTable = tableInfo{
	TableName: "skills",
	DataArr: []string{
		"C",
		"Python/Django/Flask",
		"C/C++",
		"C#",
		"HTML/CSS/JS",
		"PHP",
		"SQL",
		"Golang",
		"Linux",
		"Windows",
		"Docker",
		"Kubernetes",
	},
}

var additionalInfoTable = tableInfo{
	TableName: "additional_info",
	DataArr: []string{
		"Main character from Shenmue.",
		"Friend of Ryo from Shenmue.",
		"Second friend of Ryo from Shenmue.",
	},
}

var projectsArr = []project{
	{
		Name:    "Freelance",
		Budget:  0,
		EndTill: "01.01.2000",
	},
	{
		Name:    "Project1",
		Budget:  1_000_000,
		EndTill: "01.04.2022",
	},
	{
		Name:    "Project2",
		Budget:  500_000,
		EndTill: "01.03.2022",
	},
	{
		Name:    "Project3",
		Budget:  100_000,
		EndTill: "01.02.2022",
	},
	{
		Name:    "Project4",
		Budget:  50_000,
		EndTill: "01.01.2022",
	},
	{
		Name:    "Project5",
		Budget:  1_000_000,
		EndTill: "15.03.2022",
	},
}

var pjtTblInfo = projectsTableInfo{
	TableInfo: tableInfo{
		TableName: "projects",
		DataArr:   nil,
	},
	ProjectsArr: projectsArr,
}

var membersTable = tableInfo{
	TableName: "members",
	DataArr: []string{
		"Ryo",
		"Ren",
		"Joy",
		"David",
		"Marta",
		"Michael",
	},
}
