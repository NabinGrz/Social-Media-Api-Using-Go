package userPostModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SocialMediaPost represents a social media post.
type SocialMediaPost struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption  string             `json:"caption,omitempty" bson:"caption,omitempty"`
	User     primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Date     time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Media    []MediaDetail      `json:"media,omitempty" bson:"media,omitempty"`
	Likes    int                `json:"like,0" bson:"like,0"`
	Shares   int                `json:"shares,0" bson:"shares,0"`
	Comments int                `json:"comments,0" bson:"comments,0"`
}

type MediaDetail struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostType string             `json:"posttype,omitempty" bson:"posttype,omitempty"`
	Url      string             `json:"url,omitempty" bson:"url,omitempty"`
}
