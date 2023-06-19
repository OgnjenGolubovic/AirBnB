package startup

import (
	"authentication_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Username: "username",
		Password: "password",
		Role:     1,
		Cancels:  0,
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f80"),
		Username: "username1",
		Password: "password1",
		Role:     0,
		Cancels:  0,
	},
	{
		Id:       getObjectId("649094a74118490e167cad27"),
		Username: "username2",
		Password: "password2",
		Role:     0,
		Cancels:  0,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
