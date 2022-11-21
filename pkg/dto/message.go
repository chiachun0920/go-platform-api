package dto

import "time"

type Message struct {
	MessageType string    `json:"messageType" bson:"message_type"`
	Text        string    `json:"text"`
	Sender      string    `json:"sender"`
	CreatedAt   time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updated_at"`
}
