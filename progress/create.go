package progress

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DoCreate(c *gin.Context) {
	var json model.Progress
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
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

	operateArr := [6]string{"", "申请", "移交", "通过", "驳回", "回复"}
	msg := fmt.Sprintf("<p>%s  %s 了该审批申请。</p>", uname, operateArr[json.Operate])
	json.Opinion = fmt.Sprint(msg, json.Opinion)

	if json.Operate >= 2 || json.Operate <= 4 {
		var ap model.Approval
		result := lib.DBClient.First(&ap, json.ApprovalId)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusBadRequest,
				"msg":    result.Error.Error(),
				"data":   []string{},
			})
			return
		}

		if json.Operate == 2 {
			type OtherBody struct {
				Next string `json:"next"`
			}
			var nextB OtherBody
			if err := c.ShouldBindBodyWith(&nextB, binding.JSON); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": http.StatusBadRequest,
					"msg":    err.Error(),
					"data":   []string{},
				})
				return
			}

			ap.Approver = nextB.Next
		}

		if json.Operate == 3 {
			ap.Status = 2
		}
		if json.Operate == 4 {
			ap.Status = 3
		}

		lib.DBClient.Save(&ap)
	}

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
