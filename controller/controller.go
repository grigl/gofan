package controller

import (
	"fmt"
	"gofan/database"

	"github.com/gin-gonic/gin"
)

func HandleGetHome(ctx *gin.Context) {
	db := database.GetDatabase()

	data, err := db.Query("SELECT * FROM pages WHERE slug = 'index' LIMIT 1")
	if err != nil {
		fmt.Println("DATA ERR", err)
	}
	var slug string
	var title string
	for data.Next() {
		data.Scan(&slug, &title)
		fmt.Println("DATA", slug, title)
	}
	data.Close()
	ctx.HTML(200, "index.html", gin.H{
		"title": title,
	})
}

func HandleGetSwap(ctx *gin.Context) {
	ctx.HTML(200, "_new.html", gin.H{})
}
