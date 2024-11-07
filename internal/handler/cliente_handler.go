package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type ClienteHandler struct {
	Repo repository.ClienteRepoInterface
}

// Novo construtor para ClienteHandler
func NewClienteHandler(repo repository.ClienteRepoInterface) *ClienteHandler {
	return &ClienteHandler{Repo: repo}
}

// CreateCliente - cria um novo cliente
func (h *ClienteHandler) CreateCliente(c *gin.Context) {
	var cliente model.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(context.Background(), cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cliente"})
		return
	}
	c.JSON(http.StatusCreated, cliente)
}

// GetCliente - obtém um cliente por ID
func (h *ClienteHandler) GetCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cliente, err := h.Repo.Read(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cliente"})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

// UpdateCliente - atualiza um cliente por ID
func (h *ClienteHandler) UpdateCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var cliente model.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cliente.ID = uint(id)

	if err := h.Repo.Update(context.Background(), cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cliente"})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

// DeleteCliente - exclui um cliente por ID
func (h *ClienteHandler) DeleteCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cliente"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cliente deleted successfully"})
}
