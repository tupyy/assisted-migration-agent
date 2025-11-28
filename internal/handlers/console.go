package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAgentStatus returns the current agent status
// (GET /agent)
func (h *Handler) GetAgentStatus(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemented"})
}

// SetAgentMode changes the agent mode
// (POST /agent)
func (h *Handler) SetAgentMode(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"msg": "not implemented"})
}
