package models

import "time"

type SubTask struct {
    Name        string    `bson:"name"`
    Description string    `bson:"description"`
    Deadline    time.Time `bson:"deadline"`
    Priority    string    `bson:"priority"`
    Status      string    `bson:"status"`
}

type Task struct {
    ID          string     `bson:"_id,omitempty"`
    Name        string     `bson:"name"`
    Description string     `bson:"description"`
    Deadline    time.Time  `bson:"deadline"`
    Priority    string     `bson:"priority"`
    Status      string     `bson:"status"`
    ProjectID   string     `bson:"project_id"`
    SubTasks    []SubTask  `bson:"subtasks"`
    CreatedAt   time.Time  `bson:"created_at"`
    UpdatedAt   time.Time  `bson:"updated_at"`
}