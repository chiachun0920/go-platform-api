package dto

type CustomerProfile struct {
	UserID      string `json:"userId" bson:"user_id"`
	DisplayName string `json:"displayName" bson:"display_name"`
	PictureURL  string `json:"pictureUrl" bson:"picture_url"`
}
