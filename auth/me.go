package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowMe(c *gin.Context) {
	session := sessions.Default(c)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "ok",
		"data": gin.H{
			"uname":  session.Get("username"),
			"ticket": session.Get("ticket"),
		},
	})
}
