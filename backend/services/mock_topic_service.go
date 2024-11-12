package services

import (
	"context"
	"sustainability-platform/models"
)

type MockTopicService struct{}

func (s *MockTopicService) FetchTopics(ctx context.Context) ([]models.Topic, error) {
	return []models.Topic{
		{ID: "1", Name: "Renewable Energy", Description: "Energy sourced from renewable resources like wind, solar, etc."},
		{ID: "2", Name: "Sustainable Agriculture", Description: "Farming practices that maintain soil health and biodiversity"},
		{ID: "3", Name: "Water Conservation", Description: "Strategies to reduce water usage across industries"},
	}, nil
}
