// Copyright 2022 The TensorChord Inc.
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

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Server struct {
	Router      *gin.Engine
	AdminRouter *gin.Engine

	client *kubernetes.Clientset
	keys   []ssh.PublicKey
}

type ServerOpt struct {
	Debug      bool
	KubeConfig string
}

func New(opt ServerOpt) (*Server, error) {
	// use the current context in kubeconfig
	k8sConfig, err := clientcmd.BuildConfigFromFlags(
		"", opt.KubeConfig)
	if err != nil {
		return nil, err
	}
	cli, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())
	admin := gin.New()
	s := &Server{
		Router:      router,
		AdminRouter: admin,
		client:      cli,
		keys:        make([]ssh.PublicKey, 0),
	}
	s.bindHandlers()
	return s, nil
}

func (s *Server) bindHandlers() {
	engine := s.Router
	engine.GET("/", handlePing)
	v1 := engine.Group("/v1")
	v1.GET("/", handlePing)
	v1.POST("/pods", s.podCreate)
	v1.POST("/keys", s.KeyCreate)
	engine.POST("/config", s.OnConfig)
	engine.POST("/pubkey", s.OnPubKey)
}
func (s *Server) Run() error {
	return s.Router.Run()
}
