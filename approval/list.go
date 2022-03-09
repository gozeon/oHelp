package approval

import (
	"fmt"
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-gonic/gin"
)

func DoList(c *gin.Context) {
	var json []model.Approval
	var total int64

	p := model.Page{
		PageNum:  1,
		PageSize: 10,
	}

	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "参数错误",
			"data":   []string{},
		})
		return
	}

	if p.PageNum <= 0 {
		p.PageNum = 1
	}

	query := lib.DBClient.Model(&model.Approval{})

	id := c.Query("id")
	if len(id) > 0 {
		query.Where("id LIKE ?", fmt.Sprint("%", id, "%"))
	}

	status := c.Query("status")
	if len(status) > 0 {
		query.Where("status = ?", status)
	}

	approver := c.Query("approver")
	if len(approver) > 0 {
		query.Where("approver = ?", approver)
	}

	createUser := c.Query("createUser")
	if len(createUser) > 0 {
		query.Where("create_user = ?", createUser)
	}

	query.Count(&total)
	query.Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("updated_at desc").Find(&json)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data": gin.H{
			"total": total,
			"list":  json,
		},
	})
}
