package zakat

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/zakat")
	group.GET("", h.list)
	// group.POST("", h.create)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list zakat" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, z)
}

// func (h *Handler) create(c *gin.Context) {
// 	var in CreateZakatInput
// 	if err := c.ShouldBindJSON(&in); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	vID, err := uuid.Parse(in.VillagerID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid villager_id"})
// 		return
// 	}
// 	z := Zakat{
// 		VillagerID:  vID,
// 		Name:        in.Name,
// 		TotalPeople: in.TotalPeople,
// 		Amount:      in.Amount,
// 		Charity:     in.Charity,
// 		Category:    in.Category,
// 	}
// 	if err := h.svc.Create(c.Request.Context(), &z); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create zakat"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, z)
// }