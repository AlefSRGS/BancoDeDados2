package handler

import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
    logger *config.Logger
)

func InitializingHandlers() {
	logger = config.GetLogger("handler")
    db = config.GetDbConnection()
}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}