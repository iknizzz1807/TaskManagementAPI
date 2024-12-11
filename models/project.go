package models

import "time"

type Project struct {
	ID          string    `bson:"_id,omitempty"`
    Name        string    `bson:"name"`
    Description string    `bson:"description"`
    CreatedAt   time.Time `bson:"created_at"`
    UpdatedAt   time.Time `bson:"updated_at"`
}