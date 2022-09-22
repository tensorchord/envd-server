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

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.containerssh.io/libcontainerssh/auth"
	"go.containerssh.io/libcontainerssh/config"
	"golang.org/x/crypto/ssh"
)

func (s *Server) OnConfig(c *gin.Context) {
	var req config.Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, err)
		return
	}

	cfg := config.AppConfig{
		Backend: "sshproxy",
		SSHProxy: config.SSHProxyConfig{
			Server: "localhost",
			Port:   22222,
		},
	}
	fingerprints := []string{"SHA256:SJDm6++T0v4k5Y7InvFJ2kMQd6ui0RTi6RwvK8g3bJI"}
	for _, k := range s.authInfo {
		fingerprints = append(
			fingerprints, ssh.FingerprintSHA256(k.PublicKey))
	}
	cfg.SSHProxy.AllowedHostKeyFingerprints = fingerprints
	res := config.ResponseBody{
		Config: cfg,
	}
	c.JSON(200, res)
}

func (s *Server) OnPubKey(c *gin.Context) {
	var req auth.PublicKeyAuthRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		c.JSON(500, err)
		return
	}
	for _, k := range s.authInfo {
		logrus.Info(k.PublicKey, k.IdentityToken)
		key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(req.PublicKey.PublicKey))
		if err != nil {
			logrus.Error(err)
			c.JSON(500, err)
			return
		}
		// https://github.com/golang/go/issues/21704#issuecomment-342760478
		if bytes.Equal(key.Marshal(), k.PublicKey.Marshal()) {
			res := auth.ResponseBody{
				Success: true,
			}
			c.JSON(200, res)
			return
		}
	}
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
