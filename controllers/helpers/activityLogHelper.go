package helpers

import (
	"VincentLimarus/log-activity/configs"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActivityLogger struct {
	collection *mongo.Collection
}

func NewActivityLogger() (*ActivityLogger, error) {
	client := configs.GetMongoDB()
	if client == nil {
		return nil, fmt.Errorf("MongoDB connection failed")
	}
	return &ActivityLogger{
		collection: client.Database("logs_db").Collection("order_logs"),
	}, nil
}

func (a *ActivityLogger) LogActivity(ctx context.Context, userEmail, message string) error {
	logData := bson.M{
		"message":    message,
		"created_by": userEmail,
		"created_at": time.Now(),
	}

	_, err := a.collection.InsertOne(ctx, logData)
	if err != nil {
		return err
	}

	return err
}

func LogOrderDeletion(userEmail string, orderID uuid.UUID) {
	logger, err := NewActivityLogger()
	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
		return
	}
	logMessage := fmt.Sprintf("Order %s deleted by %s at %s", orderID, userEmail, time.Now().Format(time.RFC3339))
	_ = logger.LogActivity(context.Background(), userEmail, logMessage)	
}
