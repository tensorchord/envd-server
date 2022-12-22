// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tensorchord/envd-server/pkg/web"
)

func (s *Server) BindHandlers() {
	engine := s.Router
	web.RegisterRoute(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := engine.Group("/api/v1")

	v1.GET("/", WrapHandler(s.handlePing))
	v1.POST("/register", WrapHandler(s.register))
	v1.POST("/login", WrapHandler(s.login))
	v1.POST("/config", WrapHandler(s.OnConfig))
	v1.POST("/pubkey", WrapHandler(s.OnPubKey))

	authorized := engine.Group("/api/v1/users")
	if s.Auth {
		authorized.Use(s.AuthMiddleware())
	} else {
		authorized.Use(s.NoAuthMiddleware())
	}

	// env
	authorized.POST("/:login_name/environments", WrapHandler(s.environmentCreate))
	authorized.GET("/:login_name/environments", WrapHandler(s.environmentList))
	authorized.GET("/:login_name/environments/:name", WrapHandler(s.environmentGet))
	authorized.DELETE("/:login_name/environments/:name", WrapHandler(s.environmentRemove))
	// image
	authorized.GET("/:login_name/images/:name", WrapHandler(s.imageGet))
	authorized.GET("/:login_name/images", WrapHandler(s.imageList))
	// key
	authorized.POST("/:login_name/keys", WrapHandler(s.keyCreate))
}

type HandlerFunc func(c *gin.Context) error

func WrapHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			var serverErr *Error
			if !errors.As(err, &serverErr) {
				serverErr = &Error{
					HTTPStatusCode: http.StatusInternalServerError,
					Err:            err,
					Message:        err.Error(),
				}
			}
			serverErr.Request = c.Request.Method + " " + c.Request.URL.String()

			if gin.Mode() == "debug" {
				logrus.Debugf("error: %+v", err)
			} else {
				// Remove detailed info when in the release mode
				serverErr.Op = ""
				serverErr.Err = nil
			}

			c.JSON(serverErr.HTTPStatusCode, serverErr)
			return
		}
	}
}
