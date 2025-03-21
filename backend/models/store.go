
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const StoreCollection = "stores"

type Store struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"userId"`
	Name        string            `bson:"name" json:"name"`
	Description string            `bson:"description" json:"description"`
	Active      bool              `bson:"active" json:"active"`
	CreatedAt   time.Time         `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time         `bson:"updated_at" json:"updatedAt"`
}
