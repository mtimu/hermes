package store

import (
	"context"
	"fmt"

	"github.com/mehditeymorian/hermes/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCollection = "room-collection"

type RoomCollection struct {
	DB *mongo.Database
}

func (r RoomCollection) Create(c context.Context, room model.Room) error {
	_, err := r.DB.Collection(roomCollection).InsertOne(c, room)
	if err != nil {
		return fmt.Errorf("failed to insert room: %w", err)
	}

	return nil
}

func (r RoomCollection) Get(c context.Context, roomID string) (*model.Room, error) {
	filter := bson.D{{Key: "id", Value: roomID}}
	result := r.DB.Collection(roomCollection).FindOne(c, filter)

	err := result.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}

	var room model.Room

	err = result.Decode(&room)
	if err != nil {
		return nil, fmt.Errorf("failed to decode room: %w", err)
	}

	return &room, nil
}

func (r RoomCollection) Del(c context.Context, roomID string) error {
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return fmt.Errorf("invalid roomID: %w", err)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	_, err = r.DB.Collection(roomCollection).DeleteOne(c, filter)
	if err != nil {
		return fmt.Errorf("failed to delete room: %w", err)
	}

	return nil
}
