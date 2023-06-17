package persistence

import (
	"context"
	"fmt"

	"user_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "users"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}
func (store *UserMongoDBStore) Get(id string) (*domain.User, error) {
	filter := bson.M{"_id": ObjectIDFromHex(id)}
	return store.filterOne(filter)
}
func (store *UserMongoDBStore) Delete(id string) error {
	filter := bson.M{"_id": ObjectIDFromHex(id)}
	fmt.Println("Filterovo je : ")
	fmt.Println(filter)
	result, err := store.users.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", id)
	}
	return nil
}

func (store *UserMongoDBStore) Update(user *domain.User) error {
	fmt.Print("user in mongodb_store: ")
	fmt.Println(user)
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"username":   user.Username,
		"password":   user.Password,
		"email":      user.Email,
		"first_name": user.Name,
		"last_name":  user.Surname,
		"address":    user.Address,
	}}
	updateResult, err := store.users.UpdateByID(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Print("updateResult: ")
	fmt.Print(updateResult)
	return nil
}
func (store *UserMongoDBStore) GetByUsername(username string) (*domain.User, error) {
	fmt.Println("in GetByUsername")
	filter := bson.M{"username": username}
	fmt.Print("filter: ")
	fmt.Println(filter)
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetByEmail(email string) (*domain.User, error) {
	fmt.Println("in GetByEmail")
	filter := bson.M{"email": email}
	fmt.Print("filter: ")
	fmt.Println(filter)
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) Insert(user *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) Cancel(id string) error {
	filter := bson.M{"_id": ObjectIDFromHex(id)}
	user, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	count := user.Cancels + 1
	update := bson.M{"$set": bson.M{
		"cancels": count,
	}}
	_, err1 := store.users.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		return err1
	}
	return nil
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}
func ObjectIDFromHex(s string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return objID
}
