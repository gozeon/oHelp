package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"oHelp/lib"
	"oHelp/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func DoLogout(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	session.Save()

	loginUrl, _ := lib.CasClient.LogoutUrlForRequest(c.Request)
	u, _ := url.Parse(loginUrl)

	reqUrl, _ := utils.RequestURL(c.Request)

	reqUrl.Path = fmt.Sprintf("/%s/login", viper.GetString("APPNAME"))

	q := u.Query()
	q.Set("service", reqUrl.String())

	cookie, _ := c.Request.Cookie("_cas_session")
	cookie.MaxAge = -1
	http.SetCookie(c.Writer, cookie)

	lib.SessionsStore.Delete(cookie.Value)
	lib.TicketStore.Delete(cookie.Value)

	u.RawQuery = q.Encode()
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusForbidden,
		"msg":    "no auth",
		"data": gin.H{
			"url": u.String(),
		},
	})
}
