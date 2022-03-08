package order

import (
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DoCreate(c *gin.Context) {
	var json model.Order
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
			"data":   []string{},
		})
		return
	}

	session := sessions.Default(c)
	json.CreateUser = session.Get("username").(string)
	json.Status = true

	result := lib.DBClient.Create(&json)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   json.ID,
	})
}
