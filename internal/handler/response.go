package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorMsg struct {
	Message string `json:"message" example:"error message"`
}

type successNewSubMsg struct {
	Id int `json:"id" example:"1"`
}

type totalCostSum struct {
	Sum int `json:"sum" example:"1000"`
}

type emptyResponse struct{}

func newSuccessResponse(c *gin.Context, data interface{}) {
	logrus.Trace("ok response, data: %v", data)
	c.JSON(http.StatusOK, data)
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Trace(fmt.Sprintf("error response with code %d, message: %s", statusCode, message))
	c.AbortWithStatusJSON(statusCode, errorMsg{message})
}
