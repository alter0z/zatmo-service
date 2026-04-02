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
	group := r.Group("/api/v1/dashboard")
	group.GET("", h.getSummaryCounts)
	group.GET("/receivers", h.getReceiverCounts)
}

func (h *Handler) getSummaryCounts(c *gin.Context) {
	villagerID := c.Query("villager_id")
	z, err := h.svc.GetSummaryCounts(c.Request.Context(), villagerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get dashboard data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   z,
		"message": "success",
	})
}

func (h *Handler) getReceiverCounts(c *gin.Context) {
	receiver, err := h.svc.GetReceiverCounts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get receiver counts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   receiver,
		"message": "success",
	})
}