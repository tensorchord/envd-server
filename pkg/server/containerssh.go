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
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.containerssh.io/libcontainerssh/auth"
	"go.containerssh.io/libcontainerssh/config"
	"golang.org/x/crypto/ssh"
)

func (s *Server) OnConfig(c *gin.Context) {
	var req config.Request
	if err := c.BindJSON(&req); err != nil {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.WriteHeader(500)
		return
	}

	cfg := config.AppConfig{}
	res := config.ResponseBody{
		Config: cfg,
	}
	c.JSON(200, res)
}

func (s *Server) OnPubKey(c *gin.Context) {
	var req auth.PublicKeyAuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Writer.WriteHeader(500)
		return
	}
	for _, k := range s.authInfo {
		key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey.PublicKey))
		if err != nil {
			fmt.Println("parseerr", err)
			c.JSON(500, err)
			return
		}
		// https://github.com/golang/go/issues/21704#issuecomment-342760478
		if bytes.Equal(key.Marshal(), k.PublicKey.Marshal()) {
			res := auth.ResponseBody{
				Success: true,
			}
			fmt.Println("Success")
			c.JSON(200, res)
			return
		}
	}
	fmt.Println("fail")
	res := auth.ResponseBody{
		Success: false,
	}
	c.JSON(200, res)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
