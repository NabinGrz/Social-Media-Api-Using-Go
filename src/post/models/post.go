package userPostModel

import (
	"time"

	userModel "github.com/NabinGrz/SocialMediaApi/src/authentication/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SocialMediaPost represents a social media post.
type SocialMediaPost struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption  string             `json:"caption,omitempty" bson:"caption,omitempty"`
	User     primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Date     time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Media    []MediaDetail      `json:"media,omitempty" bson:"media,omitempty"`
	LikeBy   []userModel.User   `json:"likeby,omitempty" bson:"likeby,omitempty"`
	Shares   []userModel.User   `json:"shares,omitempty" bson:"shares,omitempty"`
	Comments []CommentDetail    `json:"comments,omitempty" bson:"comments,omitempty"`
	UserData []userModel.User   `json:"userdata,omitempty" bson:"userdata,omitempty"`
}
type CommentDetail struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Comment      string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Date         time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	User         primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	CommentUsers userModel.User     `json:"commentUsers,omitempty" bson:"commentUsers,omitempty"`
}
type MediaDetail struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostType string             `json:"posttype,omitempty" bson:"posttype,omitempty"`
	Url      string             `json:"url,omitempty" bson:"url,omitempty"`
}
