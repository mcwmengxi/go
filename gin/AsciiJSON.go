package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		ids := c.QueryMap("ids")
		page := c.DefaultQuery("page", "0")
		username := c.DefaultPostForm("username", "none")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值
		names := c.PostFormMap("names")
		// var cv map[string]interface{}
		// c.BindJSON(cv)
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"ids":      ids,
			"page":     page,
			"username": username,
			"password": password,
			// "cv": cv,
			"names": names,
			"path":  c.FullPath(),
		})
	})
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
