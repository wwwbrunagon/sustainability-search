package services

import (
	"context"
	"sustainability-platform/models"
)

type TopicService interface {
	FetchTopics(ctx context.Context) ([]models.Topic, error)
}
