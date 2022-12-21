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
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "github.com/tensorchord/envd-server/pkg/docs"
	"github.com/tensorchord/envd-server/pkg/query"
	"github.com/tensorchord/envd-server/pkg/runtime"
	runtimek8s "github.com/tensorchord/envd-server/pkg/runtime/kubernetes"
	"github.com/tensorchord/envd-server/pkg/service/image"
	"github.com/tensorchord/envd-server/pkg/service/user"
	"github.com/tensorchord/envd-server/pkg/web"
)

type Server struct {
	Router      *gin.Engine
	AdminRouter *gin.Engine
	Runtime     runtime.Provisioner

	serverFingerPrints []string

	// Auth shows if the auth is enabled.
	Auth bool

	UserService  user.Service
	ImageService image.Service
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
	conn, err := pgxpool.Connect(context.Background(), opt.DBURL)
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
		logrus.Debug("Allow CORS")
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders: []string{"*"},
		}))
	}
	admin := gin.New()

	userService := user.NewService(queries, opt.JWTSecret, opt.JWTExpirationTimeout)
	imageService := image.NewService(queries)
	s := &Server{
		Router:             router,
		AdminRouter:        admin,
		serverFingerPrints: make([]string, 0),
		Runtime:            runtimek8s.NewProvisioner(cli),
		UserService:        userService,
		ImageService:       imageService,
		Auth:               !opt.NoAuth,
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

	// Bind the HTTP handlers.
	s.BindHandlers()
	return s, nil
}

func (s *Server) BindHandlers() {
	engine := s.Router
	web.RegisterRoute(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := engine.Group("/api/v1")

	v1.GET("/", WrapHandler(s.handlePing))
	v1.POST("/register", WrapHandler(s.register))
	v1.POST("/login", WrapHandler(s.login))
	v1.POST("/config", WrapHandler(s.OnConfig))
	v1.POST("/pubkey", WrapHandler(s.OnPubKey))

	authorized := engine.Group("/api/v1/users")
	if s.Auth {
		authorized.Use(s.AuthMiddleware())
	} else {
		authorized.Use(s.NoAuthMiddleware())
	}

	// env
	authorized.POST("/:login_name/environments", WrapHandler(s.environmentCreate))
	authorized.GET("/:login_name/environments", WrapHandler(s.environmentList))
	authorized.GET("/:login_name/environments/:name", WrapHandler(s.environmentGet))
	authorized.DELETE("/:login_name/environments/:name", WrapHandler(s.environmentRemove))
	// image
	authorized.GET("/:login_name/images/:name", WrapHandler(s.imageGet))
	authorized.GET("/:login_name/images", WrapHandler(s.imageList))
}

func (s *Server) Run() error {
	return s.Router.Run()
}
