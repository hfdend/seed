package main

import (
	"log"
	"seed/cli"
	"seed/handler/api"
	"seed/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cli.Init()
	models.AutoMigrate()
	e := gin.Default()
	group := e.Group("api")
	group.GET("schemes/list", api.Schemes.List)
	group.GET("schemes/detail", api.Schemes.GetByID)
	group.POST("schemes/detail/save", api.Schemes.SaveDetail)
	group.POST("schemes/save", api.Schemes.Save)
	log.Fatalln(e.Run(":3011"))
}
