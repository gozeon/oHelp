package order

import (
	"fmt"
	"net/http"
	"oHelp/lib"
	"oHelp/model"

	"github.com/gin-gonic/gin"
)

func DoList(c *gin.Context) {
	var json []model.Order
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

	query := lib.DBClient.Model(&model.Order{})

	id := c.Query("id")
	if len(id) > 0 {
		query.Where("id LIKE ?", fmt.Sprint("%", id, "%"))
	}

	status := c.Query("status")
	if len(status) > 0 {
		query.Where("status = ?", status)
	}

	query.Count(&total)
	query.Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("updated_at desc").Find(&json)

	// result := query.Find(&json)
	// if result.Error != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": http.StatusBadRequest,
	// 		"msg":    result.Error.Error(),
	// 		"data":   []string{},
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data": gin.H{
			"total": total,
			"list":  json,
		},
	})
}
