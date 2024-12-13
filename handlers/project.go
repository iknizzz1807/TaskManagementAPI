package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/models"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projects, err := models.FetchProjects()
	if err != nil {
		http.Error(w, "Error fetching projects", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"tasks":  projects,
	})
}