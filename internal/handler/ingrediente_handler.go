package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type IngredienteHandler struct {
	Repo repository.IngredienteRepoInterface
}

// Novo construtor para IngredienteHandler
func NewIngredienteHandler(repo repository.IngredienteRepoInterface) *IngredienteHandler {
	return &IngredienteHandler{Repo: repo}
}

// CreateIngrediente - cria um novo ingrediente
func (h *IngredienteHandler) CreateIngrediente(c *gin.Context) {
	var ingrediente model.Ingrediente
	if err := c.ShouldBindJSON(&ingrediente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(context.Background(), ingrediente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ingrediente"})
		return
	}
	c.JSON(http.StatusCreated, ingrediente)
}

// GetIngrediente - obtém um ingrediente por ID
func (h *IngredienteHandler) GetIngrediente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ingrediente, err := h.Repo.Read(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ingrediente"})
		return
	}
	c.JSON(http.StatusOK, ingrediente)
}

// UpdateIngrediente - atualiza um ingrediente por ID
func (h *IngredienteHandler) UpdateIngrediente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var ingrediente model.Ingrediente
	if err := c.ShouldBindJSON(&ingrediente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ingrediente.ID = uint(id)

	if err := h.Repo.Update(context.Background(), ingrediente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ingrediente"})
		return
	}
	c.JSON(http.StatusOK, ingrediente)
}

// DeleteIngrediente - exclui um ingrediente por ID
func (h *IngredienteHandler) DeleteIngrediente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ingrediente"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ingrediente deleted successfully"})
}

