// https://golang.org/doc/articles/wiki/
package main

import (
	"art/web/itbay/src/controller"
	"log"
	"net/http"
	"os"
)

func initHandlers(serverMux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("src/assets"))
	serverMux.Handle("/src/assets/", http.StripPrefix("/src/assets/", fileServer))

	serverMux.HandleFunc("/favicon.png", controller.FaviconHandler)
	serverMux.HandleFunc("/", controller.IndexHandler)
	serverMux.HandleFunc("/member", controller.MemberPageHandler)
	serverMux.HandleFunc("/projects", controller.GetAllProjects)
	serverMux.HandleFunc("/project", controller.ProjectPageHandler)
	serverMux.HandleFunc("/skills", controller.GetAllSkills)
	serverMux.HandleFunc("/skill", controller.GetSkill)
	serverMux.HandleFunc("/about", controller.AboutInfo)
	serverMux.HandleFunc("/snake", controller.Snake)
}

func main() {
	var serverMux *http.ServeMux
	serverMux = http.NewServeMux()
	initHandlers(serverMux)

	_ = os.Setenv("PORT", "3000")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), serverMux))
}
