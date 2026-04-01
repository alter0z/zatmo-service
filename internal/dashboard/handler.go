package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/dashboard")
	group.GET("", h.getSummaryCounts)
}

func (h *Handler) getSummaryCounts(c *gin.Context) {
	villagerID := c.Query("villager_id")
	z, err := h.svc.GetSummaryCounts(c.Request.Context(), villagerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get dashboard data" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, z)
}