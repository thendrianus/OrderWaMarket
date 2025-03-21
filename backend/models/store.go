
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreDetails struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerID        primitive.ObjectID `bson:"owner_id" json:"ownerId"`
	Name           string            `bson:"name" json:"name"`
	Description    string            `bson:"description" json:"description"`
	Logo           string            `bson:"logo" json:"logo"`
	Location       string            `bson:"location" json:"location"`
	WhatsappNumber string            `bson:"whatsapp_number" json:"whatsappNumber"`
	BusinessHours  string            `bson:"business_hours" json:"businessHours"`
	Tags           []string          `bson:"tags" json:"tags"`
	Active         bool              `bson:"active" json:"active"`
	CreatedAt      time.Time         `bson:"created_at" json:"createdAt"`
	UpdatedAt      time.Time         `bson:"updated_at" json:"updatedAt"`
}
