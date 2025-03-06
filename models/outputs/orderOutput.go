package outputs

import "VincentLimarus/log-activity/models/responses"

type OrderOutput struct {
	BaseOutput
	Data responses.OrderResponseDTO `json:"data"`
}
