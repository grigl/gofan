package controller

import (
	"github.com/gin-gonic/gin"
)

func HandleGetHome(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"title": "TUTLY",
	})
}

func HandleGetSwap(ctx *gin.Context) {
	ctx.HTML(200, "_new.html", gin.H{})
}
