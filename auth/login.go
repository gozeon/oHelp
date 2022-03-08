package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func DoLogin(c *gin.Context) {
	session := sessions.Default(c)

	session.Set("ticket", c.Query("ticket"))
	session.Save()

	// back to FE page
	c.Redirect(http.StatusFound, viper.GetString("CLIENT_URL"))
}
