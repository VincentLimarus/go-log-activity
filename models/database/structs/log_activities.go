package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogActivity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Action    string             `bson:"action" json:"action"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
}
