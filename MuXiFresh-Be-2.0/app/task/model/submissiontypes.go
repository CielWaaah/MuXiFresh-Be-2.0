package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	AssignmentID primitive.ObjectID `bson:"assignment_id,omitempty" json:"assignment_id,omitempty"`
	Urls         []string           `bson:"urls,omitempty" json:"urls,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	UpdateAt     time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
