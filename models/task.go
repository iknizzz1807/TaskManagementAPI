package models

import (
	"context"
	"time"

	"errors"

	"github.com/iknizzz1807/TaskManagementAPI/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	ParentID    *primitive.ObjectID `bson:"parent_id,omitempty"` // Nil for main tasks, set for subtasks
	ProjectID   primitive.ObjectID  `bson:"project_id"`
	Name        string              `bson:"name"`
	Description string              `bson:"description"`
	Deadline    time.Time           `bson:"deadline"`
	Priority    string              `bson:"priority"`
	Status      string              `bson:"status"`
	CreatedAt   time.Time           `bson:"created_at"`
	UpdatedAt   time.Time           `bson:"updated_at"`
}

func (task *Task) ValidateParentID() error {
	if task.ParentID != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Check if the parent task exists or is a main task
		var parentTask Task
		err := utils.Db.Collection("tasks").FindOne(ctx, bson.M{"_id": task.ParentID}).Decode(&parentTask)
		if err != nil {
			return errors.New("invalid parent ID: parent task does not exist or is not a main task")
		}
	}
	return nil
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

func CreateTask(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate the ParentID
	if err := task.ValidateParentID(); err != nil {
		return err
	}

	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := utils.Db.Collection("tasks").InsertOne(ctx, task)
	return err
}

func DeleteTask(taskID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete the main task
	result, err := utils.Db.Collection("tasks").DeleteOne(ctx, bson.M{"_id": taskID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	// Delete subtasks
	_, err = utils.Db.Collection("tasks").DeleteMany(ctx, bson.M{"parent_id": taskID})
	return err
}
