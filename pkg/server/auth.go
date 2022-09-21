// Copyright 2022 TensorChord Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"

	"github.com/tensorchord/envd-server/api/types"
)

type AuthInfo struct {
	IdentityToken string
	PublicKey     ssh.PublicKey
}

func (s *Server) auth(c *gin.Context) {
	var req types.AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, err)
		return
	}

	key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey))
	if err != nil {
		c.JSON(500, err)
		return
	}
	s.authInfo = append(s.authInfo, AuthInfo{
		PublicKey:     key,
		IdentityToken: req.IdentityToken,
	})
	res := types.AuthResponse{
		IdentityToken: req.IdentityToken,
		Status:        "login succeeded",
	}
	c.JSON(200, res)
}
