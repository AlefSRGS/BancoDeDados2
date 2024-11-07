package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type VendaHandler struct {
	Repo repository.VendaRepoInterface
}

// Novo construtor para VendaHandler
func NewVendaHandler(repo repository.VendaRepoInterface) *VendaHandler {
	return &VendaHandler{Repo: repo}
}

// CreateVenda - cria uma nova venda
func (h *VendaHandler) CreateVenda(c *gin.Context) {
	var venda model.Venda
	if err := c.ShouldBindJSON(&venda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(context.Background(), venda); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create venda"})
		return
	}
	c.JSON(http.StatusCreated, venda)
}

// GetVenda - obtém uma venda por ID
func (h *VendaHandler) GetVenda(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	venda, err := h.Repo.Read(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get venda"})
		return
	}
	c.JSON(http.StatusOK, venda)
}

// UpdateVenda - atualiza uma venda por ID
func (h *VendaHandler) UpdateVenda(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var venda model.Venda
	if err := c.ShouldBindJSON(&venda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	venda.ID = uint(id)

	if err := h.Repo.Update(context.Background(), venda); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update venda"})
		return
	}
	c.JSON(http.StatusOK, venda)
}

// DeleteVenda - exclui uma venda por ID
func (h *VendaHandler) DeleteVenda(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Repo.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete venda"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Venda deleted successfully"})
}
