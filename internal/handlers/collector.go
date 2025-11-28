package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCollectorStatus returns the collector status
// (GET /collector)
func (h *Handler) GetCollectorStatus(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemented"})
}

// StartCollector starts inventory collection
// (POST /collector)
func (h *Handler) StartCollector(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemented"})
}

// GetInventory returns the collected inventory
// (GET /collector/inventory)
func (h *Handler) GetInventory(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemented"})
}

// StopCollector stops the collection
// (DELETE /collector)
func (h *Handler) StopCollector(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemented"})
}
