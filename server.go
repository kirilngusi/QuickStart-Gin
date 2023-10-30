package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kirilngusi/QuickStart-Gin/controller"
	"github.com/kirilngusi/QuickStart-Gin/middlewares"
	"github.com/kirilngusi/QuickStart-Gin/service"
	"io"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func SetupLogOutput() {
	f, _ := os.Create("kiril.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := server.Run(":" + port)
	if err != nil {
		return
	}
}
