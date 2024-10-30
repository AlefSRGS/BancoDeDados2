package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/handler"
)

func initRoutes(router *gin.Engine){
	crud := router.Group("/crud")

	crud.GET("/hello", handler.HelloCRUD)
}