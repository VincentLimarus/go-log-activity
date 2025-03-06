package responses

import (
	"time"

	"github.com/google/uuid"
)

type OrderResponseDTO struct {
	ID          uuid.UUID `json:"id" db:"id"`
	OrderStatus string    `json:"order_status" db:"order_status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
