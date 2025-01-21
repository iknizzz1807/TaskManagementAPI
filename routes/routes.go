package routes

import (
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/handlers"
	"github.com/iknizzz1807/TaskManagementAPI/middleware"
)

func SetupRoutes() {
	http.Handle("/projects", middleware.CORSMiddleware(middleware.AuthMiddleware(middleware.CacheMiddleware(http.HandlerFunc(handlers.GetProjects)))))
	http.Handle("/project/{id}", middleware.CORSMiddleware(middleware.AuthMiddleware(middleware.CacheMiddleware(http.HandlerFunc(handlers.FetchTasksInProject)))))
	http.Handle("/project", middleware.CORSMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateProject))))
	http.Handle("/project/{id}", middleware.CORSMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteProject))))
	http.Handle("/project/{id}", middleware.CORSMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProject))))

	http.Handle("/tasks", middleware.CORSMiddleware(middleware.AuthMiddleware(middleware.CacheMiddleware(http.HandlerFunc(handlers.GetTasks)))))
	http.Handle("/task", middleware.CORSMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateTask))))
	http.Handle("/task/{id}", middleware.CORSMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteTask))))
	http.Handle("/task/{id}", middleware.CORSMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateTask))))

	http.Handle("/register", middleware.CORSMiddleware(http.HandlerFunc(handlers.RegisterUser)))
	http.Handle("/login", middleware.CORSMiddleware(http.HandlerFunc(handlers.LoginUser)))
}
