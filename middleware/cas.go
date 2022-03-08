package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"oHelp/lib"
	"oHelp/utils"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/cas.v2"
)

func CasAuth(exclude ...string) gin.HandlerFunc {
	handler := lib.CasClient.HandleFunc(func(writer http.ResponseWriter, request *http.Request) {
	})
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
		session := sessions.Default(c)

		rPath := c.Request.URL.Path

		if len(exclude) > 0 {
			if utils.Contains(exclude, path.Base(rPath)) {
				c.Next()
				return
			}
		}

		if cas.IsAuthenticated(c.Request) && session.Get("ticket") != nil {
			session.Set("username", cas.Username(c.Request))
			session.Save()
			c.Next()
		} else {
			loginUrl, _ := lib.CasClient.LoginUrlForRequest(c.Request)
			u, _ := url.Parse(loginUrl)

			reqUrl, _ := utils.RequestURL(c.Request)

			reqUrl.Path = fmt.Sprintf("/%s/login", viper.GetString("APPNAME"))

			q := u.Query()
			q.Set("service", reqUrl.String())

			u.RawQuery = q.Encode()

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusForbidden,
				"msg":    "no auth",
				"data": gin.H{
					"url": u.String(),
				},
			})
			c.Abort()
		}
	}
}
