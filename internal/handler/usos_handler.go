package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type UsosHandler struct {
	Repo repository.UsosRepoInterface
}

// Novo construtor para UsosHandler
func NewUsosHandler(repo repository.UsosRepoInterface) *UsosHandler {
	return &UsosHandler{Repo: repo}
}

// CreateUsos - cria um novo uso de ingrediente em um prato
func (h *UsosHandler) CreateUsos(c *gin.Context) {
	var uso model.Usos
	if err := c.ShouldBindJSON(&uso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(context.Background(), uso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uso"})
		return
	}
	c.JSON(http.StatusCreated, uso)
}

// GetUsos - obtém um uso por IDs de prato e ingrediente
func (h *UsosHandler) GetUsos(c *gin.Context) {
	idPrato, err := strconv.Atoi(c.Param("id_prato"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prato ID"})
		return
	}

	idIngrediente, err := strconv.Atoi(c.Param("id_ingrediente"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ingrediente ID"})
		return
	}

	uso, err := h.Repo.Read(context.Background(), idPrato, idIngrediente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get uso"})
		return
	}
	c.JSON(http.StatusOK, uso)
}

// DeleteUsos - exclui um uso por IDs de prato e ingrediente
func (h *UsosHandler) DeleteUsos(c *gin.Context) {
	idPrato, err := strconv.Atoi(c.Param("id_prato"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prato ID"})
		return
	}

	idIngrediente, err := strconv.Atoi(c.Param("id_ingrediente"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ingrediente ID"})
		return
	}

	if err := h.Repo.Delete(context.Background(), idPrato, idIngrediente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete uso"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Uso deleted successfully"})
}
