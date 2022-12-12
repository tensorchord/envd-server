// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"k8s.io/apimachinery/pkg/runtime"
	kubernetes "k8s.io/client-go/kubernetes/fake"

	runtimek8s "github.com/tensorchord/envd-server/pkg/runtime/kubernetes"
	"github.com/tensorchord/envd-server/pkg/server"
)

func NewServer(objects ...runtime.Object) (*server.Server, error) {
	cli := kubernetes.NewSimpleClientset(objects...)

	router := gin.New()
	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
	admin := gin.New()
	s := &server.Server{
		Router:      router,
		AdminRouter: admin,
		Runtime:     runtimek8s.NewProvisioner(cli),
		Auth:        false,
	}

	s.BindHandlers()
	return s, nil
}
