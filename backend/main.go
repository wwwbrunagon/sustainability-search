package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sustainability-platform/handlers"
	"sustainability-platform/middleware" // Import the middleware package
	"sustainability-platform/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up MongoDB connection
	var topicService services.TopicService
	useMock := os.Getenv("USE_MOCK_SERVICE") // set this environment variable to "true" for mock service

	if useMock == "true" {
		topicService = &services.MockTopicService{}
	} else {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}
		defer client.Disconnect(context.Background())
		topicCollection := client.Database("sustainability").Collection("topics")
		topicService = &services.MongoTopicService{Collection: topicCollection}
	}

	// Initialize handler
	topicHandler := handlers.TopicHandler{Service: topicService}

	// Set up routes with CORS middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/api/topics", topicHandler.GetTopics)

	// Wrap the router with CORS middleware
	corsHandler := middleware.EnableCORS(mux)

	// Start the server with CORS enabled
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
