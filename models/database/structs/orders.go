package structs

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          int       `db:"id" json:"id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	OrderStatus string    `db:"order_status" json:"order_status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
