package controller

import (
	"art/web/itbay/src/db/read"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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

func AboutInfo(resWriter http.ResponseWriter, request *http.Request) {
	htmlTpl := template.Must(template.ParseFiles("src/view/about.html"))
	_ = htmlTpl.Execute(resWriter, nil)
}

func FaviconHandler(resWriter http.ResponseWriter, request *http.Request) {

}
