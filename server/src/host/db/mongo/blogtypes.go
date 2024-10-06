package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	Title   string   `bson:"Title" json:"Title"`
	Tags    []string `bson:"Tags" json:"Tags"`
	Brief   string   `bson:"Brief" json:"Brief"`
	Content string   `bson:"Content" json:"Content"`
}
