package services

import (
	"VincentLimarus/log-activity/controllers/helpers"
	"VincentLimarus/log-activity/models/outputs"
	"VincentLimarus/log-activity/models/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context) {
	var DeleteOrderRequestDTO requests.DeleteOrderRequestDTO

	if err := c.ShouldBindJSON(&DeleteOrderRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteOrder(c, DeleteOrderRequestDTO)
	c.JSON(code, output)
}

func AuthOrderService(router *gin.RouterGroup) {
	router.POST("/order/delete", DeleteOrder)
}
