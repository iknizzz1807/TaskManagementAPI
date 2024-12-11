package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iknizzz1807/TaskManagementAPI/models"
	"github.com/iknizzz1807/TaskManagementAPI/utils"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
    // Set headers
    w.Header().Set("Content-Type", "application/json")

    // Get database connection
    db := utils.GetDB() // Todo: Avoid being called multiple times
    if db == nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }

    // Create context with timeout
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    // Initialize projects slice
    var projects []models.Project

    // Execute find query with context
    cursor, err := db.Collection("projects").Find(ctx, map[string]interface{}{})
    if err != nil {
        http.Error(w, "Error fetching projects", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    // Decode results
    if err := cursor.All(ctx, &projects); err != nil {
        http.Error(w, "Error decoding projects", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":   "success",
        "projects": projects,
    })
}