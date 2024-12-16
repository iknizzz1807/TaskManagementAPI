package routes

import (
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/handlers"
)

func SetupRoutes() {
	http.HandleFunc("GET /projects", handlers.GetProjects)
	http.HandleFunc("GET /project/{id}", handlers.FetchTasksInProject)
	http.HandleFunc("POST /project", handlers.CreateProject)
	http.HandleFunc("DELETE /project/{id}", handlers.DeleteProject)
	http.HandleFunc("PUT /project/{id}", handlers.UpdateProject)

	http.HandleFunc("GET /tasks", handlers.GetTasks)
	http.HandleFunc("POST /task", handlers.CreateTask)
	http.HandleFunc("DELETE /task/{id}", handlers.DeleteTask)
	http.HandleFunc("PUT /task/{id}", handlers.UpdateTask)
}
