package items

import (
	"net/http"
	services "reuse-api/services/item"
	"strconv"

	dto "reuse-api/dto/items"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service services.ItemService
}

func NewItemHandler(service services.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var req dto.ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDAny, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userID := uint(userIDAny.(float64))
	item, err := h.service.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := dto.ItemResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Category:    item.Category,
		Condition:   item.Condition,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		OwnerID:     item.OwnerID,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *ItemHandler) GetAllItems(c *gin.Context) {
	items, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) GetItemByID(c *gin.Context) {
	idParam := c.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	item, err := h.service.GetById(uint(parsedID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
		return
	}
	c.JSON(http.StatusOK, item)
}
