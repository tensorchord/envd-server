//go:build !debug

package web

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/envd-server/dashboard"
)

func RegisterRoute(route *gin.Engine) {
	webRoot, _ := fs.Sub(dashboard.DistFS, "dist")
	route.StaticFS("/dashboard", http.FS(webRoot))
}
