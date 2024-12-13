package models

import (
	"context"
	"time"

	"github.com/iknizzz1807/TaskManagementAPI/utils"
)

type Project struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func FetchProjects() ([]Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var projects []Project
	cursor, err := utils.Db.Collection("projects").Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
