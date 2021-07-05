package controller

import (
	"art/web/itbay/src/db/create"
	"art/web/itbay/src/db/del"
	"art/web/itbay/src/db/read"
	"art/web/itbay/src/db/update"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getTimeStr(timeStr string) string {
	tillTime, _ := time.Parse("2006-01-02", timeStr)
	tillTimeStr := tillTime.Format("02.01.2006")
	return tillTimeStr
}

func IndexHandler(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/index.html"))
	_ = htmlTpl.Execute(resWriter, read.GetAllMembers())
}

func MemberPageHandler(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/member.html"))
	id, _ := strconv.Atoi(request.FormValue("member_id"))
	log.Printf("Member id=%d\n", id)
	_ = htmlTpl.Execute(resWriter, read.GetMember(id))
}

func GetAllProjects(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/projects.html"))
	_ = htmlTpl.Execute(resWriter, read.GetAllProjects())
}

func AddProject(resWriter http.ResponseWriter, request *http.Request) {
	pjtName := request.FormValue("project_name")
	pjtBgt := request.FormValue("project_budget")
	tillTimeStr := getTimeStr(request.FormValue("end_till"))
	if err := create.AddProject(pjtName, pjtBgt, tillTimeStr); err != nil {
		log.Printf("%s <%s>", err, pjtName)
	}
	http.Redirect(resWriter, request, "/projects", http.StatusFound)
}

func DelPjt(resWriter http.ResponseWriter, request *http.Request) {
	pjtId := request.FormValue("pjt_id")
	if err := del.DeleteProject(pjtId); err != nil {
		log.Printf("%s <%s>", err, pjtId)
	}
	http.Redirect(resWriter, request, "/projects", http.StatusFound)
}

func UpdPjt(resWriter http.ResponseWriter, request *http.Request) {
	pjtId := request.FormValue("project_id")
	newPjtName := request.FormValue("project_name")
	newPjtBgt := request.FormValue("project_budget")
	newTillTimeStr := getTimeStr(request.FormValue("end_till"))
	if err := update.UpdPjt(pjtId, newPjtName, newPjtBgt, newTillTimeStr); err != nil {
		log.Printf("%s <%s>", err, newPjtName)
	}
	http.Redirect(resWriter, request, "/projects", http.StatusFound)
}

func ProjectPageHandler(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/project.html"))
	id, _ := strconv.Atoi(request.FormValue("project_id"))
	log.Printf("Project id=%d\n", id)
	_ = htmlTpl.Execute(resWriter, read.GetProject(id))
}

func GetAllSkills(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/skills.html"))
	_ = htmlTpl.Execute(resWriter, read.GetAllSkills())
}

func GetSkill(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/skill.html"))
	id, _ := strconv.Atoi(request.FormValue("skill_id"))
	_ = htmlTpl.Execute(resWriter, read.GetSkill(id))
}

func AddSkill(resWriter http.ResponseWriter, request *http.Request) {
	skillName := request.FormValue("signature")
	if err := create.AddSkill(skillName); err != nil {
		log.Printf("%s <%s>", err, skillName)
	}
	http.Redirect(resWriter, request, "/skills", http.StatusFound)
}

func DelSkill(resWriter http.ResponseWriter, request *http.Request) {
	skillId := request.FormValue("skill_id")
	i, _ := strconv.Atoi(skillId)
	if err := del.DeleteSkill(i); err != nil {
		log.Printf("%s <%d>", err, i)
	}
	http.Redirect(resWriter, request, "/skills", http.StatusFound)
}

func UpdSkill(resWriter http.ResponseWriter, request *http.Request) {
	skillId := request.FormValue("skill_id")
	newSkillName := request.FormValue("new_skill_name")
	if err := update.UpdSkill(skillId, newSkillName); err != nil {
		log.Printf("%s <%s>", err, skillId)
	}
	http.Redirect(resWriter, request, "/skills", http.StatusFound)
}

func AboutInfo(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/about.html"))
	_ = htmlTpl.Execute(resWriter, nil)
}

func Snake(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/snake.html"))
	_ = htmlTpl.Execute(resWriter, nil)
}

func FaviconHandler(resWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(resWriter, request, "favicon.png")
}
