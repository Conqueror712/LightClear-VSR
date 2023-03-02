package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/thinkerou/favicon"
)

// Tips: Get the video from user in Front-end
// Use FFmpeg to change the video code
// Output to the database and back to the Front-end to let it have a download buttom

// For more, it can deploy on the cloud_server and we can search on internet
// Finally, is the UI design and make the web more beautiful

type Video struct {
	gorm.Model
	Name string
	Path string
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:BqV?eGcc_1o+@/ffmpeg-list?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	db.AutoMigrate(&Video{})

	// 初始化Gin
	router := gin.Default()

	router.Use(favicon.New("static/images/lightclear-logo.ico")) // 这里如果添加了东西然后再运行没有变化，请重启浏览器，浏览器有缓存

	// 加载静态页面
	router.LoadHTMLGlob("templates/*") // 一种是全局加载，一种是加载指定的文件

	// 加载资源文件
	router.Static("/static", "./static")

	// 将根路径与 index.html 文件进行绑定
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This data is come from Go background.",
		})
	})

	// 上传视频文件的路由处理函数
	router.POST("/upload", func(c *gin.Context) {
		// 从HTTP请求中读取上传的文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 保存上传的文件到本地磁盘
		path := fmt.Sprintf("uploads/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 转码视频文件
		output := fmt.Sprintf("output_%s", file.Filename)
		cmd := exec.Command("ffmpeg", "-i", path, "-vcodec", "libx264", "-acodec", "aac", "-strict", "-2", output)
		if err := cmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 将转码后的文件保存到本地磁盘
		outputPath := fmt.Sprintf("outputs/%s", output)
		if err := exec.Command("mv", output, outputPath).Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 将视频信息保存到数据库中
		video := Video{Name: output, Path: outputPath}
		if err := db.Create(&video).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Upload and transcode video successfully!",
		})
	})

	// 启动HTTP服务器
	router.Run(":8080")
}
