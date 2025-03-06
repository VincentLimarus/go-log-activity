package helpers

import (
	"VincentLimarus/log-activity/configs"
	"VincentLimarus/log-activity/models/outputs"
	"VincentLimarus/log-activity/models/requests"
	"VincentLimarus/log-activity/models/responses"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteOrder(c *gin.Context, DeleteOrderRequestDTO requests.DeleteOrderRequestDTO) (int, any) {
	db := configs.GetDB()
	mongoClient := configs.GetMongoDB()

	if db == nil {
		return 500, outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Database connection failed",
		}
	}

	id := DeleteOrderRequestDTO.ID

	if id == uuid.Nil {
		return 400, outputs.BadRequestOutput{
			Code:    400,
			Message: "Invalid order ID",
		}
	}

	query := "SELECT id, order_status, created_at FROM orders WHERE id = $1"
	var order responses.OrderResponseDTO

	err := db.Get(&order, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 404, outputs.NotFoundOutput{
				Code:    404,
				Message: "Order not found",
			}
		}
		return 500, outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	email, exists := c.Get("user_email")
	if !exists {
		return 500, outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	logMessage := fmt.Sprintf("Order %s deleted by %s at %s", order.ID, email.(string), time.Now().Format(time.RFC3339))
	logCollection := mongoClient.Database("logs_db").Collection("order_logs")

	logData := bson.M{
		"message":    logMessage,
		"created_by": email.(string),
		"created_at": time.Now(),
	}

	_, err = logCollection.InsertOne(context.Background(), logData)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Failed to log activity to MongoDB",
		}
	}

	deleteQuery := "DELETE FROM orders WHERE id = $1"
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

    output := outputs.OrderOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.OrderResponseDTO{
		ID : order.ID,
		OrderStatus : order.OrderStatus,
		CreatedAt : order.CreatedAt,
	}
    return 200, output
}