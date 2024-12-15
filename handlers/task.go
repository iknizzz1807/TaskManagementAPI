package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iknizzz1807/TaskManagementAPI/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract query parameters
	name := r.URL.Query().Get("name")
	status := r.URL.Query().Get("status")
	priority := r.URL.Query().Get("priority")

	tasks, err := models.FetchTasks(name, status, priority)
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"tasks":  tasks,
	})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Change from string to objectID
	var task struct {
		models.Task
		ParentIDStr  string `json:"parent_id"`
		ProjectIDStr string `json:"project_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if task.ParentIDStr != "" {
		parentID, err := primitive.ObjectIDFromHex(task.ParentIDStr)
		if err != nil {
			http.Error(w, "Invalid parent ID", http.StatusBadRequest)
			return
		}
		task.Task.ParentID = &parentID
	}

	if task.ProjectIDStr != "" {
		projectID, err := primitive.ObjectIDFromHex(task.ProjectIDStr)
		if err != nil {
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}
		task.Task.ProjectID = projectID
	} else {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	if err := models.CreateTask(&task.Task); err != nil {
		if err.Error() == "internal error" {
			http.Error(w, "Error creating task", http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"task":   task.Task,
	})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteTask(taskID); err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting task", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskIDStr := r.PathValue("id")
	if taskIDStr == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	taskID, err := primitive.ObjectIDFromHex(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Change from string to objectID
	var task struct {
		models.Task
		ParentIDStr  string `json:"parent_id"`
		ProjectIDStr string `json:"project_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if task.ParentIDStr != "" {
		parentID, err := primitive.ObjectIDFromHex(task.ParentIDStr)
		if err != nil {
			http.Error(w, "Invalid parent ID", http.StatusBadRequest)
			return
		}
		task.Task.ParentID = &parentID
	}

	if task.ProjectIDStr != "" {
		projectID, err := primitive.ObjectIDFromHex(task.ProjectIDStr)
		if err != nil {
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}
		task.Task.ProjectID = projectID
	} else {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	if err := models.UpdateTask(taskID, &task.Task); err != nil {
		if err.Error() == "internal error" {
			http.Error(w, "Error updating task", http.StatusInternalServerError)
		} else if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"task":   task.Task,
	})
}
