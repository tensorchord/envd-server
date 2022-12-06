// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/pkg/service/user"
)

func (s *Server) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		amURI := AuthMiddlewareURIRequest{}
		if err := c.BindUri(&amURI); err != nil {
			respondWithError(c, http.StatusUnauthorized,
				fmt.Sprintf("auth failed: %v", err))
			c.Next()
			return
		}

		logrus.WithFields(logrus.Fields{
			"login-name-in-uri": amURI.LoginName,
		}).Debug("debug")

		amr := AuthMiddlewareHeaderRequest{}
		if err := c.BindHeader(&amr); err != nil {
			respondWithError(c, http.StatusUnauthorized,
				fmt.Sprintf("auth failed: %v", err))
			c.Next()
			return
		}

		userService := user.NewService(s.Queries, s.JWTSecret, s.JWTExpirationTimeout)
		loginName, err := userService.ValidateJWT(amr.JWTToken)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized,
				fmt.Sprintf("failed to validate the JWT: %v", err))
			c.Next()
			return
		}
		if loginName != amURI.LoginName {
			logrus.WithFields(logrus.Fields{
				"login-name":        loginName,
				"login-name-in-uri": amURI.LoginName,
			}).Debug("login name in JWT does not match the login name in URI")
			respondWithError(c, http.StatusUnauthorized,
				"loginname mismatch")
			c.Next()
			return
		}
		c.Set(ContextLoginName, loginName)
		c.Next()
	}
}

// NoAuthMiddleware is a middleware that does not auth the user.
func (s *Server) NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		amURI := AuthMiddlewareURIRequest{}
		if err := c.BindUri(&amURI); err != nil {
			respondWithError(c, http.StatusUnauthorized,
				fmt.Sprintf("auth failed: %v", err))
			c.Next()
			return
		}

		c.Set(ContextLoginName, amURI.LoginName)
		c.Next()
	}
}

// nolint:unparam
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
