package main

import (
	"log"
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/routes"
	"github.com/iknizzz1807/TaskManagementAPI/utils"
)

func main() {
	utils.ConnectDB()
	routes.SetupRoutes()

	log.Println("Server starting on PORT 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}