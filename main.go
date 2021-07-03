// https://golang.org/doc/articles/wiki/
package main

import (
	"art/web/itbay/src/controller"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func initHandlers(serverMux *http.ServeMux) {
	// Передача файловому серверу пути, по которому находятся все статические файлы.
	fileServer := http.FileServer(http.Dir("src/assets"))
	// Регистрация патерна /assets/ в http multiplexer.
	serverMux.Handle("/src/assets/", http.StripPrefix("/src/assets/", fileServer))

	serverMux.HandleFunc("/", controller.IndexHandler)
	serverMux.HandleFunc("/member", controller.MemberPageHandler)
	serverMux.HandleFunc("/projects", controller.GetAllProjects)
	serverMux.HandleFunc("/project", controller.ProjectPageHandler)
	serverMux.HandleFunc("/skills", controller.GetAllSkills)
	serverMux.HandleFunc("/skill", controller.GetSkill)
	serverMux.HandleFunc("/about", controller.AboutInfo)

	serverMux.HandleFunc("/favicon.ico", controller.FaviconHandler)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var serverMux *http.ServeMux
	serverMux = http.NewServeMux()
	initHandlers(serverMux)

	/*DB JOB! -----------------------------------------------------------------*/
	//db.CreateTables()
	//db.CreateEntries()
	/*DB JOB! -----------------------------------------------------------------*/

	log.Fatal(http.ListenAndServe(":"+port, serverMux))
}
