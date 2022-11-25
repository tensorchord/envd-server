//go:build debug

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) {
	// Do nothing, we don't need dashboard in this case
	route.GET("/dashboard", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, dashboard!")
	})
}
