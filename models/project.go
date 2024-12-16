package models

import (
	"context"
	"errors"
	"time"

	"github.com/iknizzz1807/TaskManagementAPI/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (project *Project) Validate() error {
	if project.Name == "" {
		return errors.New("name is required")
	}
	if project.Description == "" {
		return errors.New("description is required")
	}
	return nil
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

func FetchProject(projectID primitive.ObjectID) (Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var project Project
	err := utils.Db.Collection("projects").FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Project{}, errors.New("project not found")
		}
		return Project{}, err
	}

	return project, nil
}

func CreateProject(project *Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate project before inserting
	if err := project.Validate(); err != nil {
		return err
	}

	project.ID = primitive.NewObjectID()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	_, err := utils.Db.Collection("projects").InsertOne(ctx, project)
	if err != nil {
		return errors.New("internal error")
	}
	return nil
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

func UpdateProject(projectID primitive.ObjectID, updatedProject *Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate the updated project
	if err := updatedProject.Validate(); err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":        updatedProject.Name,
			"description": updatedProject.Description,
			"updated_at":  updatedProject.UpdatedAt,
		},
	}

	result, err := utils.Db.Collection("projects").UpdateOne(ctx, bson.M{"_id": projectID}, update)
	if err != nil {
		return errors.New("internal error")
	}

	if result.MatchedCount == 0 {
		return errors.New("project not found")
	}

	return nil
}
