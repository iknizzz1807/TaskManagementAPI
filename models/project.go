package models

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/TaskManagementAPI/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
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

func CreateProject(project *Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project.ID = primitive.NewObjectID()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	_, err := utils.Db.Collection("projects").InsertOne(ctx, project)
	return err
}

func DeleteProject(projectID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := utils.Db.Collection("projects").DeleteOne(ctx, bson.M{"_id": projectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("project not found")
	}

	// Delete associated tasks
	_, err = utils.Db.Collection("tasks").DeleteMany(ctx, bson.M{"project_id": projectID})
	return err
}
