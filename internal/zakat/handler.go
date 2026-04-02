package zakat

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/api/v1/zakat")
	group.GET("", h.list)
	group.GET("/villager", h.getVillager)
	group.POST("", h.create)
	group.PUT("/:id", h.update)
	group.DELETE("/:id", h.delete)
}

type ListQuery struct {
    VillagerID string `form:"villager_id"`
    Name       string `form:"name"`
    Category   *bool  `form:"category"`
    Date       string `form:"date"`
}

func (h *Handler) list(c *gin.Context) {
	var q ListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	z, err := h.svc.List(
		c.Request.Context(),
		q.VillagerID,
		q.Name,
		q.Category,
		datePtr,
	)
	if z == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list zakat"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   z,
		"message": "success",
	})
}
func (h *Handler) getVillager(c *gin.Context) {
	v, err := h.svc.GetVillager(c.Request.Context())
	if v == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get villager"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   v,
		"message": "success",
	})
}

func (h *Handler) create(c *gin.Context) {
	var in CreateZakatInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// vID, err := uuid.Parse(in.VillagerID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid villager_id"})
	// 	return
	// }
	z := Zakat{
		VillagerID:  in.VillagerID,
		Name:        in.Name,
		TotalPeople: in.TotalPeople,
		Charity:     in.Charity,
		Category:    in.Category,
	}
	z, err := h.svc.Create(c.Request.Context(), in)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create zakat"})
    return
  }
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   z,
		"message": "success",
	})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")

	var in UpdateZakatInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	in.ID = ID

	z, err := h.svc.Update(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update zakat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   z,
		"message": "success",
	})
}

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete zakat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   nil,
		"message": "success",
	})
}