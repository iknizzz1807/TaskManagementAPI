package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iknizzz1807/TaskManagementAPI/models"
	"github.com/iknizzz1807/TaskManagementAPI/utils"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    db := utils.GetDB() // Todo: Avoid being called multiple times
    if db == nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }

    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    var tasks []models.Task
    
    cursor, err := db.Collection("tasks").Find(ctx, map[string]interface{}{})
    if err != nil {
        http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    if err := cursor.All(ctx, &tasks); err != nil {
        http.Error(w, "Error decoding tasks", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "success",
        "tasks":  tasks,
    })
}