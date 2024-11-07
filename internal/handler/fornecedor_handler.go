package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type FornecedorHandler struct {
	Repo repository.FornecedorRepoInterface
}

// Novo construtor para FornecedorHandler
func NewFornecedorHandler(repo repository.FornecedorRepoInterface) *FornecedorHandler {
	return &FornecedorHandler{Repo: repo}
}

// CreateFornecedor - cria um novo fornecedor
func (h *FornecedorHandler) CreateFornecedor(c *gin.Context) {
	var fornecedor model.Fornecedor
	if err := c.ShouldBindJSON(&fornecedor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(context.Background(), fornecedor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create fornecedor"})
		return
	}
	c.JSON(http.StatusCreated, fornecedor)
}

// GetFornecedor - obtém um fornecedor por ID
func (h *FornecedorHandler) GetFornecedor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	fornecedor, err := h.Repo.Read(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fornecedor"})
		return
	}
	c.JSON(http.StatusOK, fornecedor)
}

// UpdateFornecedor - atualiza um fornecedor por ID
func (h *FornecedorHandler) UpdateFornecedor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var fornecedor model.Fornecedor
	if err := c.ShouldBindJSON(&fornecedor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fornecedor.ID = uint(id)

	if err := h.Repo.Update(context.Background(), fornecedor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update fornecedor"})
		return
	}
	c.JSON(http.StatusOK, fornecedor)
}

// DeleteFornecedor - exclui um fornecedor por ID
func (h *FornecedorHandler) DeleteFornecedor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete fornecedor"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fornecedor deleted successfully"})
}
