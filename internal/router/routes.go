package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/handler"
)


func initRoutes(router *gin.Engine){
	handler.InitializingHandlers()
	
	// cliente := router.Group("/cliente")
	// fornecedor := router.Group("/fornecedor")
	// ingrediente := router.Group("/ingrediente")
	// prato := router.Group("/prato")
	// venda := router.Group("/venda")
	router.GET("/ping", handler.Ping)
}