package helpers

import (
	"VincentLimarus/log-activity/configs"
	"VincentLimarus/log-activity/models/outputs"
	"VincentLimarus/log-activity/models/requests"
	"VincentLimarus/log-activity/models/responses"
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateDummyOrders() {
	db := configs.GetDB()
	if db == nil {
		log.Println("Database connection failed")
		return
	}

	orderStatuses := []string{"Pending", "Completed", "Cancelled", "Processing", "Failed"}

	for range 10 {
		order := responses.OrderResponseDTO{
			ID:          uuid.New(),
			OrderStatus: orderStatuses[rand.Intn(len(orderStatuses))],
			CreatedAt:   time.Now(),
		}

		_, err := db.Exec(
			"INSERT INTO orders (id, order_status, created_at) VALUES ($1, $2, $3)",
			order.ID, order.OrderStatus, order.CreatedAt,
		)
		if err != nil {
			log.Println("Failed to insert order:", err)
		} else {
			log.Printf("Inserted order: %s - %s", order.ID, order.OrderStatus)
		}
	}
}


func DeleteOrder(c *gin.Context, DeleteOrderRequestDTO requests.DeleteOrderRequestDTO) (int, any) {
	db := configs.GetDB()
	if db == nil {
		return 500, outputs.InternalServerErrorOutput{Code: 500, Message: "Database connection failed"}
	}

	id := DeleteOrderRequestDTO.ID
	if id == uuid.Nil {
		return 400, outputs.BadRequestOutput{Code: 400, Message: "Invalid, UUID Not Found"}
	}

	query := "SELECT id, order_status, created_at FROM orders WHERE id = $1"
	var order responses.OrderResponseDTO
	err := db.Get(&order, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 404, outputs.NotFoundOutput{Code: 404, Message: "Order not found"}
		}
		return 500, outputs.InternalServerErrorOutput{Code: 500, Message: "Internal Server Error"}
	}

	email, exists := c.Get("user_email")
	if !exists {
		return 500, outputs.InternalServerErrorOutput{Code: 500, Message: "Internal Server Error"}
	}

	go LogOrderDeletion(email.(string), order.ID)

	deleteQuery := "DELETE FROM orders WHERE id = $1"
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Code: 500, Message: "Internal Server Error"}
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