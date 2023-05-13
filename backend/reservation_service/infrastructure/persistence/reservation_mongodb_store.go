package persistence

import (
	"context"

	"reservation_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "reservations"
	COLLECTION = "reservation"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) Get(id string) (*domain.AccommodationReservation, error) {
	filter := bson.M{"_id": ObjectIDFromHex(id)}
	return store.filterOne(filter)
}

func (store *ReservationMongoDBStore) Insert(reservation *domain.AccommodationReservation) error {
	result, err := store.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) AccommodationReservationRequest(reservation *domain.AccommodationReservation) error {
	result, err := store.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) DeleteAll() {
	store.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*domain.AccommodationReservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ReservationMongoDBStore) filterOne(filter interface{}) (reservation *domain.AccommodationReservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&reservation)
	return
}

func decode(cursor *mongo.Cursor) (reservations []*domain.AccommodationReservation, err error) {
	for cursor.Next(context.TODO()) {
		var reservation domain.AccommodationReservation
		err = cursor.Decode(&reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &reservation)
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
