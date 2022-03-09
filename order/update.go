package order

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-gonic/gin"
)

func DoUpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var oldD model.Order

	result := lib.DBClient.First(&oldD, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    result.Error.Error(),
			"data":   []string{},
		})
		return
	}

	var newD model.Order
	if err := c.ShouldBindJSON(&newD); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
			"data":   []string{},
		})
		return
	}

	oldD.Status = newD.Status

	session := sessions.Default(c)
	uname := session.Get("username").(string)
	desc := "<p> %s  <strong><span style='color: %s;'>%s</span>  </strong>了工单。</p>"

	if oldD.Status {
		desc = fmt.Sprintf(desc, uname, "#47b881", "重新打开")

	} else {
		desc = fmt.Sprintf(desc, uname, "#ec4c47", "关闭")

	}

	lib.DBClient.Create(&model.Comment{
		Description: desc,
		CreateUser:  uname,
		OrderId:     oldD.ID,
	})

	lib.DBClient.Save(oldD)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   oldD,
	})
}
