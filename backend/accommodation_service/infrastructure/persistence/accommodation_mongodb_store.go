package persistence

import (
	"context"

	"accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accommodations"
	COLLECTION = "accommodation"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) GetAll() ([]*domain.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) GetAllByHost(id string) ([]*domain.Accommodation, error) {
	filter := bson.M{"hostId": ObjectIDFromHex(id)}
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) UpdatePrice(accommodation *domain.Accommodation) error {
	filter := bson.M{"_id": accommodation.Id}
	update := bson.M{"$set": bson.M{
		"price":      accommodation.Price,
		"isPerGuest": accommodation.IsPerGuest,
		"hasWeekend": accommodation.HasWeekend,
		"hasSummer":  accommodation.HasSummer,
	}}

	_, err := store.accommodations.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (store *AccommodationMongoDBStore) AddFreeDates(accommodation *domain.Accommodation) error {
	filter := bson.M{"_id": accommodation.Id}
	update := bson.M{"$set": bson.M{
		"dates": accommodation.Dates,
	}}

	_, err := store.accommodations.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (store *AccommodationMongoDBStore) Insert(accommodation *domain.Accommodation) error {
	result, err := store.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return err
	}
	accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccommodationMongoDBStore) DeleteAll() {
	store.accommodations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccommodationMongoDBStore) DeleteAccomodations(id string) {
	store.accommodations.DeleteMany(context.TODO(), bson.M{"hostId": ObjectIDFromHex(id)})
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (accommodation *domain.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&accommodation)
	return
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var accommodation domain.Accommodation
		err = cursor.Decode(&accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &accommodation)
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
