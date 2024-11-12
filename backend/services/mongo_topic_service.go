package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sustainability-platform/models"
)

type MongoTopicService struct {
	Collection *mongo.Collection
}

func (s *MongoTopicService) FetchTopics(ctx context.Context) ([]models.Topic, error) {
	cursor, err := s.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var topics []models.Topic
	for cursor.Next(ctx) {
		var topic models.Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	return topics, nil
}
