package main

import (
	"fmt"
	"log"
	"net/http"
	"oHelp/auth"
	"oHelp/lib"
	"oHelp/middleware"
	"oHelp/model"
	"oHelp/order"
	"oHelp/upload"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	lib.InitCas()
	lib.InitDb()

	lib.DBClient.AutoMigrate(&model.Order{})

	err := os.MkdirAll(filepath.Join(".", "static"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	appName := viper.GetString("APPNAME")
	port := fmt.Sprintf(":%s", viper.Get("PORT"))

	db, _ := lib.DBClient.DB()
	defer db.Close()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	store := cookie.NewStore([]byte(appName))
	r.Use(sessions.Sessions(fmt.Sprint("_", appName, "_session"), store))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	oHelp := r.Group(fmt.Sprintf("/%s", appName))
	oHelp.Static("/static", "./static")
	oHelp.Use(middleware.CasAuth("login", "logout"))
	{
		oHelp.GET("/login", auth.DoLogin)
		oHelp.GET("/logout", auth.DoLogout)
		oHelp.GET("/me", auth.ShowMe)
		oHelp.POST("/upload", upload.SingleUpload)

		orderR := oHelp.Group("/order")
		{
			orderR.GET("/", order.DoList)
			orderR.POST("/", order.DoCreate)
		}
	}
	r.Run(port)
}