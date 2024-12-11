package routes

import (
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/controllers"
)

func SetupRoutes() {
    http.HandleFunc("GET /projects", controllers.GetProjects)
	http.HandleFunc("GET /tasks", controllers.GetTasks)
}