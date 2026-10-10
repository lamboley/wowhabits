// Package api is the HTTP REST adapter.
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter returns the HTTP handler serving the REST API.
func NewRouter() http.Handler {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	return r
}
