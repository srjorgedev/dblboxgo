package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjorgedev/dblboxgo/internal/domain/data"
)

type Handler struct {
	repo data.Repository
}

func NewDataHandler(repo data.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetTags(c *gin.Context) {
	u, err := h.repo.GetCachedTags()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetChapters(c *gin.Context) {
	u, err := h.repo.GetCachedChapters()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetRarities(c *gin.Context) {
	u, err := h.repo.GetCachedRarities()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetTypes(c *gin.Context) {
	u, err := h.repo.GetCachedTypes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetAffinities(c *gin.Context) {
	u, err := h.repo.GetCachedAffinities()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}