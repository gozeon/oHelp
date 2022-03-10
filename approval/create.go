package approval

import (
	"fmt"
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
	uname := session.Get("username").(string)
	json.CreateUser = uname

	json.Status = 1

	result := lib.DBClient.Create(&json)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    result.Error.Error(),
			"data":   []string{},
		})
		return
	}

	lib.DBClient.Create(&model.Progress{
		Operate:    1,
		Opinion:    fmt.Sprint(uname, " 发起申请。"),
		CreateUser: uname,
		ApprovalId: json.ID,
	})

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   json.ID,
	})
}
