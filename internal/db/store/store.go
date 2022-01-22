package store

import "go.mongodb.org/mongo-driver/mongo"

type Store struct {
	RoomCollection RoomCollection
}

func New(db *mongo.Database) Store {
	return Store{RoomCollection: RoomCollection{DB: db}}
}
