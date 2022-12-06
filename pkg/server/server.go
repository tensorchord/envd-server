// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "github.com/tensorchord/envd-server/pkg/docs"
	"github.com/tensorchord/envd-server/pkg/query"
	"github.com/tensorchord/envd-server/pkg/util"
	"github.com/tensorchord/envd-server/pkg/web"
)

type Server struct {
	Router      *gin.Engine
	AdminRouter *gin.Engine
	Queries     *query.Queries
	Client      kubernetes.Interface

	serverFingerPrints []string

	// Auth shows if the auth is enabled.
	Auth bool
	// JWTSecret is the secret used to sign the JWT token.
	JWTSecret string
	// JWTExpirationTimeout is the expiration time of the JWT token.
	JWTExpirationTimeout time.Duration
}

type Opt struct {
	Debug       bool
	KubeConfig  string
	HostKeyPath string
	DBURL       string

	NoAuth               bool
	JWTSecret            string
	JWTExpirationTimeout time.Duration
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
	conn, err := pgx.Connect(context.Background(), opt.DBURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = util.ApplySchema(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to apply schema: %v\n", err)
		os.Exit(1)
	}

	queries := query.New(conn)

	router := gin.New()
	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
		}))
	}
	admin := gin.New()
	s := &Server{
		Router:             router,
		AdminRouter:        admin,
		Client:             cli,
		Queries:            queries,
		serverFingerPrints: make([]string, 0),
	}

	// Load host key.
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

	// Set the auth information.
	s.Auth = !opt.NoAuth
	s.JWTSecret = opt.JWTSecret
	s.JWTExpirationTimeout = opt.JWTExpirationTimeout

	// Bind the HTTP handlers.
	s.BindHandlers()
	return s, nil
}

func (s *Server) BindHandlers() {
	engine := s.Router
	web.RegisterRoute(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := engine.Group("/api/v1")

	v1.GET("/", s.handlePing)
	v1.POST("/register", s.register)
	v1.POST("/login", s.login)
	v1.POST("/config", s.OnConfig)
	v1.POST("/pubkey", s.OnPubKey)

	authorized := engine.Group("/api/v1/users")
	if s.Auth {
		authorized.Use(s.AuthMiddleware())
	} else {
		authorized.Use(s.NoAuthMiddleware())
	}

	// env
	authorized.POST("/:login_name/environments", s.environmentCreate)
	authorized.GET("/:login_name/environments", s.environmentList)
	authorized.GET("/:login_name/environments/:name", s.environmentGet)
	authorized.DELETE("/:login_name/environments/:name", s.environmentRemove)
	// image
	authorized.GET("/:login_name/images/:name", s.imageGet)
	authorized.GET("/:login_name/images", s.imageList)
}

func (s *Server) Run() error {
	return s.Router.Run()
}
