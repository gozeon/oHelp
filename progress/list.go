package progress

import (
	"fmt"
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-gonic/gin"
)

func DoListByApprovalId(c *gin.Context) {
	id := c.Param("id")

	var json []model.Progress

	query := lib.DBClient.Model(&model.Progress{})

	operate := c.Query("operate")
	if len(operate) > 0 {
		query.Where("operate = ?", operate)
	}

	updated_at_sort := c.DefaultQuery("updated_at", "desc")

	query.Where("approval_id = ?", id).Order(fmt.Sprint("updated_at ", updated_at_sort)).Find(&json)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   json,
	})
}
