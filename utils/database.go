package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"time"

	"github.com/iknizzz1807/TaskManagementAPI/models"
)

var (
    client  *mongo.Client
    Db      *mongo.Database
)

func ConnectDB() {
    var err error
    client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
        return
    }
    Db = client.Database("task_management")
}

func FetchTasks() ([]models.Task, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var tasks []models.Task
    cursor, err := Db.Collection("tasks").Find(ctx, map[string]interface{}{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    if err := cursor.All(ctx, &tasks); err != nil {
        return nil, err
    }

    return tasks, nil
}

func FetchProjects() ([]models.Project, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var projects []models.Project
    cursor, err := Db.Collection("projects").Find(ctx, map[string]interface{}{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    if err := cursor.All(ctx, &projects); err != nil {
        return nil, err
    }

    return projects, nil
}