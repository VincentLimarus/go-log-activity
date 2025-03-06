package outputs

type BaseOutput struct {
	Code    int    `json:"code" default:"200"`
	Message string `json:"message" default:"Success: {message}"`
}

type BadRequestOutput struct {
	Code    int    `json:"code" default:"400"`
	Message string `json:"message" default:"Bad Request: {message}"`
}

type InternalServerErrorOutput struct {
	Code    int    `json:"code" default:"500"`
	Message string `json:"message" default:"Internal Server Error: {message}"`
}

type UnauthorizedOutput struct {
	Code    int    `json:"code" default:"401"`
	Message string `json:"message" default:"Unauthorized: {message}"`
}

type NotFoundOutput struct {
	Code    int    `json:"code" default:"404"`
	Message string `json:"message" default:"Not Found: {message}"`
}