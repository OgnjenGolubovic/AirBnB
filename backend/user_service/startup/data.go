package startup

import (
	"user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Username: "username",
		Password: "password",
		Role:     1,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}