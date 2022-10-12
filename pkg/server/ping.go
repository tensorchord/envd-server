// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handlePing is the handler for ping requrets.
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (s *Server) handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world!", "goto": "https://github.com/tensorchord/envd"})
}
