package main

import (
	"gofan/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", controller.HandleGetHome)
	router.GET("/swap", controller.HandleGetSwap)

	router.Run()
}
