package services

import (
	"VincentLimarus/log-activity/controllers/helpers"
	"VincentLimarus/log-activity/models/outputs"
	"VincentLimarus/log-activity/models/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var RegisterUserRequestDTO requests.RegisterUserRequestDTO

	if err := c.ShouldBindJSON(&RegisterUserRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.RegisterUser(RegisterUserRequestDTO)
	c.JSON(code, output)
}

func LoginUser(c *gin.Context){
	var LoginUserRequestDTO requests.LoginUserRequestDTO

	if err := c.ShouldBindJSON(&LoginUserRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code , output, token := helpers.LoginUser(LoginUserRequestDTO)
	c.SetCookie("Authorization", token, 3600*12, "/", "localhost", true, true)
	c.SetSameSite(http.StatusOK)
	c.JSON(code, output)
}

func BaseUserService(router *gin.RouterGroup) {
	router.POST("/user/register", RegisterUser)
	router.POST("/user/login", LoginUser)
}