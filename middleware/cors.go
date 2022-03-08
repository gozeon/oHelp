package middleware

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CORSMiddleware() gin.HandlerFunc {
	clientUrl := viper.GetString("CLIENT_URL")
	u, _ := url.Parse(clientUrl)

	return cors.New(cors.Config{
		AllowOrigins:     []string{fmt.Sprint("http://", u.Host), fmt.Sprint("https://", u.Host)},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
