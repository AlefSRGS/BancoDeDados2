package handler

import (
	"net/http"

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

func HelloCRUD(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello-CRUD",
	})
}