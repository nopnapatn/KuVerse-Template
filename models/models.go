package models

import (
	"time"
)

type Report struct {
	UserUid   string                 `json:"user_uid" bson:"user_uid"`
	VerseUid  string                 `json:"verse_uid" bson:"verse_uid"`
	Type      string                 `json:"type" bson:"type"`
	Note      string                 `json:"note" bson:"note"`
	Data      map[string]interface{} `json:"data" bson:"data"`
	CreatedAt time.Time              `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time             `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time             `json:"deleted_at" bson:"deleted_at"`
}
