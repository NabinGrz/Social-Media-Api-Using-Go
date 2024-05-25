package userPostModel

type UserPost struct {
	UserID string `json:"user,omitempty" bson:"user,omitempty"`
	PostID string `json:"posts,omitempty" bson:"posts,omitempty"`
}
