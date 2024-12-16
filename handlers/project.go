package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		"count":  len(projects),
		"tasks":  projects,
	})
}

func FetchTasksInProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectIDStr := r.PathValue("id")

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var project models.Project

	project, err = models.FetchProject(projectID)
	if err != nil {
		if err.Error() == "project not found" {
			http.Error(w, "Project not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching project", http.StatusInternalServerError)
		}
		return
	}

	tasks, err := models.FetchTasksByProjectID(projectID)
	if err != nil {
		http.Error(w, "Error fetching tasks in project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"project": project,
		"tasks":   tasks,
		"count":   len(tasks),
	})
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := models.CreateProject(&project); err != nil {
		if err.Error() == "internal error" {
			http.Error(w, "Error creating project", http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"project": project,
	})
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	projectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteProject(projectID); err != nil {
		if err.Error() == "project not found" {
			http.Error(w, "Project not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting project", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectIDStr := r.PathValue("id")
	if projectIDStr == "" {
		http.Error(w, "Missing project ID", http.StatusBadRequest)
		return
	}

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	project.ID = projectID

	if err := models.UpdateProject(projectID, &project); err != nil {
		if err.Error() == "internal error" {
			http.Error(w, "Error updating project", http.StatusInternalServerError)
		} else if err.Error() == "project not found" {
			http.Error(w, "Project not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"project": project,
	})
}
