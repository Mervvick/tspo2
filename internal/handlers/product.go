// internal/handlers/product.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"digital-market/internal/models"
	"digital-market/internal/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Create(c *gin.Context) {
	// Проверка прав администратора
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}
	
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.productService.Create(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	product, err := h.productService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	// Проверка прав администратора
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}
	
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	product.ID = id
	if err := h.productService.Update(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	// Проверка прав администратора
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}
	
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	if err := h.productService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandler) List(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	categoryIDStr := c.Query("category_id")
	
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}
	
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	
	var categoryID *uuid.UUID
	if categoryIDStr != "" {
		id, err := uuid.Parse(categoryIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
		categoryID = &id
	}
	
	products, err := h.productService.List(c.Request.Context(), offset, limit, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, products)
}

// Вспомогательная функция для проверки прав администратора
func isAdmin(c *gin.Context) bool {
	value, exists := c.Get("is_admin")
	if !exists {
		return false
	}
	isAdmin, ok := value.(bool)
	return ok && isAdmin
}
