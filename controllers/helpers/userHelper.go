package helpers

import (
	"VincentLimarus/log-activity/configs"
	"VincentLimarus/log-activity/models/outputs"
	"VincentLimarus/log-activity/models/requests"
	"VincentLimarus/log-activity/models/responses"
	"VincentLimarus/log-activity/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func RegisterUser(RegisterUserRequestDTO requests.RegisterUserRequestDTO) (int, interface{}) {
	db := configs.GetDB()

	userID := uuid.New()

	isActive := true
	if RegisterUserRequestDTO.IsActive != nil {
		isActive = *RegisterUserRequestDTO.IsActive
	}

	createdAt := time.Now()
	if RegisterUserRequestDTO.CreatedAt != nil {
		createdAt = *RegisterUserRequestDTO.CreatedAt
	}

	updatedAt := time.Now()
	if RegisterUserRequestDTO.UpdatedAt != nil {
		updatedAt = *RegisterUserRequestDTO.UpdatedAt
	}

	query := `
		INSERT INTO users (id, name, email, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, email, is_active, created_at, updated_at
	`

	var user responses.UserResponseDTO
	err := db.QueryRowx(query, userID, RegisterUserRequestDTO.Name, RegisterUserRequestDTO.Email,
		isActive, createdAt, updatedAt,
	).Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	output := outputs.RegisterUserOutputDTO{}
	output.Code = 200
	output.Message = "Success: User has been registered"
	output.Data = user

	return 200, output
}



func LoginUser(LoginUserRequestDTO requests.LoginUserRequestDTO) (int, interface{}, string) {
	db := configs.GetDB()

	if LoginUserRequestDTO.Email == "" {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: Email is required",
		}
		return 400, output, "Error"
	}

	query := `
		SELECT id, name, email, is_active, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user responses.UserResponseDTO
	err := db.QueryRowx(query, LoginUserRequestDTO.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			output := outputs.BadRequestOutput{
				Code:    404,
				Message: "User not found",
			}
			return 404, output, "User Error"
		}

		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output, "Internal Error"
	}

	token, tokenErr := utils.CreateJWTToken(user.ID, user.Email)

	if tokenErr != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: fmt.Sprintf("Internal Server Error: %v", tokenErr),
		}
		return 500, output, tokenErr.Error()
	}


	output := outputs.LoginUserOutputDTO{}
	output.Code = 200
	output.Message = "Success: User has been logged in"
	output.Data = user

	return 200, output, token
}