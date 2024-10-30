package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloCRUD(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello-CRUD",
	})
}