// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *Server) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		amURI := AuthMiddlewareURIRequest{}
		if err := c.BindUri(&amURI); err != nil {
			respondWithError(c, NewError(http.StatusUnauthorized, err, "auth.middleware.bind-uri"))
			c.Next()
			return
		}

		logrus.WithFields(logrus.Fields{
			"login-name-in-uri": amURI.LoginName,
		}).Debug("debug")

		amr := AuthMiddlewareHeaderRequest{}
		if err := c.BindHeader(&amr); err != nil {
			respondWithError(c, NewError(http.StatusUnauthorized, err, "auth.middleware"))
			c.Next()
			return
		}

		loginName, err := s.UserService.ValidateJWT(amr.JWTToken)
		if err != nil {
			respondWithError(c, NewError(http.StatusUnauthorized, err, "user.validateJWT"))
			c.Next()
			return
		}
		if loginName != amURI.LoginName {
			logrus.WithFields(logrus.Fields{
				"login-name":        loginName,
				"login-name-in-uri": amURI.LoginName,
			}).Debug("login name in JWT does not match the login name in URI")
			respondWithError(c, NewError(http.StatusUnauthorized, err, "user.validateJWT"))
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
			respondWithError(c, NewError(http.StatusUnauthorized, err, "auth.middleware.bind-uri"))
			c.Next()
			return
		}

		c.Set(ContextLoginName, amURI.LoginName)
		c.Next()
	}
}

// nolint:unparam
func respondWithError(c *gin.Context, err error) {
	var serverErr *Error
	if errors.As(err, &serverErr) {
		c.AbortWithStatusJSON(serverErr.HTTPStatusCode, serverErr)
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, err)
}
