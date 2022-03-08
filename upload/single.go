package upload

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SingleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusForbidden,
			"msg":    err.Error(),
			"data":   "",
		})
		return
	}

	filename := filepath.Base(file.Filename)
	newFilepath := fmt.Sprint("./static/", time.Now().Unix(), "-", filename)

	if err := c.SaveUploadedFile(file, newFilepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
			"data":   "",
		})
		return
	}

	u, _ := url.Parse(viper.GetString("STATIC_URL"))
	u.Path = path.Join(u.Path, filepath.Base(newFilepath))

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "",
		"data": gin.H{
			"url": u.String(),
		},
	})

}
