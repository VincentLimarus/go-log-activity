package outputs

import "VincentLimarus/log-activity/models/responses"

type RegisterUserOutputDTO struct {
	BaseOutput
	Data responses.UserResponseDTO `json:"data"`
}

type LoginUserOutputDTO struct {
	BaseOutput
	Data responses.UserResponseDTO `json:"data"`
}

