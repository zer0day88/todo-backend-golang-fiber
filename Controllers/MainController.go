package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type main_resp struct {
	Message string `json:"message"`
}

func GetMain(c *gin.Context) {
	var data main_resp
	data.Message = "Welcome to API TODO"
	c.JSON(http.StatusOK, data)
}
