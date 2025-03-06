package requests

import "time"

type RegisterUserRequestDTO struct {
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	IsActive  *bool      `json:"is_active,omitempty"`   
	CreatedAt *time.Time `json:"created_at,omitempty"`  
	UpdatedAt *time.Time `json:"updated_at,omitempty"`  
}

type LoginUserRequestDTO struct {
	Email string `json:"email" binding:"required"`
}
