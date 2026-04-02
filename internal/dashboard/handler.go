package dashboard

import (
	"net/http"
	"time"

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

type ListQuery struct {
    VillagerID string `form:"villager_id"`
    Date       string `form:"date"`
}

func (h *Handler) getSummaryCounts(c *gin.Context) {
	var q ListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var datePtr *time.Time
	if q.Date != "" {
		t, err := time.Parse("2006-01-02", q.Date)
		if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
				return
		}
		datePtr = &t
	}

	z, err := h.svc.GetSummaryCounts(c.Request.Context(), q.VillagerID, datePtr)
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