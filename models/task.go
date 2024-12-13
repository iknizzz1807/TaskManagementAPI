package models

import (
	"context"
	"time"

	"github.com/iknizzz1807/TaskManagementAPI/utils"
)

type SubTask struct {
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Deadline    time.Time `bson:"deadline"`
	Priority    string    `bson:"priority"`
	Status      string    `bson:"status"`
}

type Task struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Deadline    time.Time `bson:"deadline"`
	Priority    string    `bson:"priority"`
	Status      string    `bson:"status"`
	ProjectID   string    `bson:"project_id"`
	SubTasks    []SubTask `bson:"subtasks"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func FetchTasks() ([]Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tasks []Task
	cursor, err := utils.Db.Collection("tasks").Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
