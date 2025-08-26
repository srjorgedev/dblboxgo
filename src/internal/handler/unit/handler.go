package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjorgedev/dblboxgo/src/internal/domain/unit"
)

type Handler struct {
	repo unit.Repository
}

func NewUnitHandler(repo unit.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetUnitByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	u, err := h.repo.GetUnitByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetAllUnitSummaries(c *gin.Context) {
	summaries, err := h.repo.GetAllUnitSummariesCached()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

func (h *Handler) GetUnitSummaryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	u, err := h.repo.GetUnitSummaryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}
