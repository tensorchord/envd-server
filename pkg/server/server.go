// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Server struct {
	Router      *gin.Engine
	AdminRouter *gin.Engine

	client   *kubernetes.Clientset
	authInfo []AuthInfo
}

type Opt struct {
	Debug      bool
	KubeConfig string
}

func New(opt Opt) (*Server, error) {
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
		authInfo:    make([]AuthInfo, 0),
	}
	s.bindHandlers()
	return s, nil
}

func (s *Server) bindHandlers() {
	engine := s.Router
	engine.GET("/", handlePing)
	v1 := engine.Group("/v1")
	v1.GET("/", handlePing)
	v1.POST("/environments", s.environmentCreate)
	v1.POST("/auth", s.auth)
	v1.POST("/config", s.OnConfig)
	v1.POST("/pubkey", s.OnPubKey)
}

func (s *Server) Run() error {
	return s.Router.Run()
}
