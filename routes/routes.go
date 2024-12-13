package routes

import (
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/handlers"
)

func SetupRoutes() {
	http.HandleFunc("GET /projects", handlers.GetProjects)
	http.HandleFunc("GET /tasks", handlers.GetTasks)
}
