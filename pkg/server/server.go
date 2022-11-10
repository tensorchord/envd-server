// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/tensorchord/envd-server/api/types"
	_ "github.com/tensorchord/envd-server/pkg/docs"
	"github.com/tensorchord/envd-server/pkg/query"
)

type Server struct {
	Router      *gin.Engine
	AdminRouter *gin.Engine
	Queries     *query.Queries
	Conn        *pgx.Conn

	client             *kubernetes.Clientset
	serverFingerPrints []string
	imageInfo          []types.ImageInfo
}

type Opt struct {
	Debug       bool
	KubeConfig  string
	HostKeyPath string
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

	// Connect to database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	queries := query.New(conn)

	router := gin.New()
	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
	admin := gin.New()
	s := &Server{
		Router:             router,
		AdminRouter:        admin,
		client:             cli,
		Conn:               conn,
		Queries:            queries,
		serverFingerPrints: make([]string, 0),
	}
	if opt.HostKeyPath != "" {
		// read private key file
		pemBytes, err := os.ReadFile(opt.HostKeyPath)
		if err != nil {
			return nil, errors.Wrapf(
				err, "reading private key %s failed", opt.HostKeyPath)
		}
		if privateKey, err := ssh.ParsePrivateKey(pemBytes); err != nil {
			return nil, err
		} else {
			logrus.Debugf("load host key from %s", opt.HostKeyPath)
			fingerPrint := ssh.FingerprintSHA256(privateKey.PublicKey())
			s.serverFingerPrints = append(s.serverFingerPrints, fingerPrint)
		}
	}
	s.bindHandlers()
	return s, nil
}

func (s *Server) bindHandlers() {
	engine := s.Router

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := engine.Group("/v1")
	authorized := engine.Group("/v1/users")
	authorized.Use(s.AuthMiddleware())
	{
		// env
		authorized.POST("/:identity_token/environments", s.environmentCreate)
		authorized.GET("/:identity_token/environments", s.environmentList)
		authorized.GET("/:identity_token/environments/:name", s.environmentGet)
		authorized.DELETE("/:identity_token/environments/:name", s.environmentRemove)
		// image
		authorized.GET("/:identity_token/images/:name", s.imageGet)
		authorized.GET("/:identity_token/images", s.imageList)
	}

	v1.GET("/", s.handlePing)
	v1.POST("/auth", s.auth)
	v1.POST("/config", s.OnConfig)
	v1.POST("/pubkey", s.OnPubKey)
}

func (s *Server) Run() error {
	return s.Router.Run()
}
