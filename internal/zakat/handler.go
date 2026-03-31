package zakat

import (
	"net/http"

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

func (h *Handler) list(c *gin.Context) {
	z, err := h.svc.List(c.Request.Context())
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