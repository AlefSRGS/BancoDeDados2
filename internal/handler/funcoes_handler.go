package handler

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
)

type FuncoesHandler struct {
	repo *repository.FuncoesRepo
}

func NewFuncoesHandler(repo *repository.FuncoesRepo) *FuncoesHandler {
	return &FuncoesHandler{repo: repo}
}

// Aplica reajuste percentual nos pratos
func (h *FuncoesHandler) Reajuste(c *gin.Context) {
	percentualStr := c.Query("percentual")
	percentual, err := strconv.ParseFloat(percentualStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid percentage"})
		return
	}

	err = h.repo.Reajuste(percentual)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply adjustment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Adjustment applied successfully"})
}

// Sorteia um cliente e adiciona pontos
func (h *FuncoesHandler) SorteioCliente(c *gin.Context) {
	err := h.repo.SorteioCliente()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute draw"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer points awarded successfully"})
}

// Retorna estatísticas de vendas
func (h *FuncoesHandler) EstatisticasVendas(c *gin.Context) {
	maisVendido, menosVendido, err := h.repo.EstatisticasVendas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve statistics"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"most_sold":    maisVendido,
		"least_sold":   menosVendido,
	})
}
