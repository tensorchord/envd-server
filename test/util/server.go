package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tensorchord/envd-server/pkg/server"
	ginlogrus "github.com/toorop/gin-logrus"
	"k8s.io/apimachinery/pkg/runtime"
	kubernetes "k8s.io/client-go/kubernetes/fake"
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
		Client:      cli,
	}

	s.BindHandlers(false)
	return s, nil
}
