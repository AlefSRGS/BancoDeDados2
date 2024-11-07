package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type PratoHandler struct {
	Repo repository.PratoRepoInterface
}

// Novo construtor para PratoHandler
func NewPratoHandler(repo repository.PratoRepoInterface) *PratoHandler {
	return &PratoHandler{Repo: repo}
}

// CreatePrato - cria um novo prato
func (h *PratoHandler) CreatePrato(c *gin.Context) {
	var prato model.Prato
	if err := c.ShouldBindJSON(&prato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(context.Background(), prato); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create prato"})
		return
	}
	c.JSON(http.StatusCreated, prato)
}

// GetPrato - obtém um prato por ID
func (h *PratoHandler) GetPrato(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	prato, err := h.Repo.Read(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prato"})
		return
	}
	c.JSON(http.StatusOK, prato)
}

// UpdatePrato - atualiza um prato por ID
func (h *PratoHandler) UpdatePrato(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var prato model.Prato
	if err := c.ShouldBindJSON(&prato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prato.ID = uint(id)

	if err := h.Repo.Update(context.Background(), prato); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update prato"})
		return
	}
	c.JSON(http.StatusOK, prato)
}

// DeletePrato - exclui um prato por ID
func (h *PratoHandler) DeletePrato(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete prato"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prato deleted successfully"})
}
