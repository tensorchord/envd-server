//go:build !debug

/*
   Copyright The TensorChord Inc.
   Copyright The BuildKit Authors.
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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
	route.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(webRoot))
	route.GET("/dashboard")
	route.NoRoute(func(c *gin.Context) {
		c.FileFromFS("/index.html", http.FS(webRoot))
	})
}
