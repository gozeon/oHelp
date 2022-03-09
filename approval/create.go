package approval

import (
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DoCreate(c *gin.Context) {
	var json model.Approval
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

	result := lib.DBClient.Create(&json)
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
		"data":   json.ID,
	})
}
