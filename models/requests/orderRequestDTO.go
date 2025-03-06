package requests

import "github.com/google/uuid"

type DeleteOrderRequestDTO struct {
	ID uuid.UUID `json:"id"`
}