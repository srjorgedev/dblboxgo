package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjorgedev/dblboxgo/internal/domain/equipment"
)

type Handler struct {
	repo equipment.Repository
}

func NewEquipmentHandler(repo equipment.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetAllEquipmentSummaries(c *gin.Context) {
	summaries, err := h.repo.GetAllEquipmentSummariesCached()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get equipment summaries"})
		return
	}
	c.JSON(http.StatusOK, summaries)
}
