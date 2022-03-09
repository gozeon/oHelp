package comment

import (
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-gonic/gin"
)

func DoListByOrderId(c *gin.Context) {
	id := c.Param("id")

	var json []model.Comment

	query := lib.DBClient.Model(&model.Comment{})

	query.Where("order_id = ?", id).Order("updated_at asc").Find(&json)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   json,
	})
}
