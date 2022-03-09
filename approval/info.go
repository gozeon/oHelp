package approval

import (
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-gonic/gin"
)

func DoInfo(c *gin.Context) {
	id := c.Param("id")

	var json model.Approval

	result := lib.DBClient.First(&json, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    result.Error.Error(),
			"data":   []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   json,
	})
}
