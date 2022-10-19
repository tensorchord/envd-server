// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tensorchord/envd-server/api/types"
)

func (s *Server) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		amr := types.AuthMiddlewareRequest{}
		if err := c.BindUri(&amr); err != nil {
			respondWithError(c, http.StatusUnauthorized,
				fmt.Sprintf("auth failed: %v", err))
		}
		for _, k := range s.authInfo {
			if k.IdentityToken == amr.IdentityToken {
				c.Set("identity_token", amr.IdentityToken)
				c.Next()
				return
			}
		}
		respondWithError(c, http.StatusUnauthorized,
			"failed to auth the identity_token")
	}
}

// nolint:unparam
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
