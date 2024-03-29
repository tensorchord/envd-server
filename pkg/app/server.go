// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package app

import (
	"time"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"

	"github.com/tensorchord/envd-server/pkg/server"
	"github.com/tensorchord/envd-server/pkg/version"
)

type EnvdServerApp struct {
	*cli.App
}

func New() EnvdServerApp {
	internalApp := cli.NewApp()
	internalApp.EnableBashCompletion = true
	internalApp.Name = "envd-server"
	internalApp.Usage = "HTTP backend server for envd"
	internalApp.HideHelpCommand = true
	internalApp.HideVersion = true
	internalApp.Version = version.GetVersion().String()
	internalApp.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output in logs",
		},
		&cli.PathFlag{
			Name:    "kubeconfig",
			Usage:   "kubeconfig path",
			EnvVars: []string{"KUBE_CONFIG", "KUBECONFIG"},
		},
		&cli.PathFlag{
			Name:    "hostkey",
			Usage:   "hostkey in the backend pod, used to generate fingerprint here",
			EnvVars: []string{"ENVD_SERVER_HOST_KEY"},
		},
		&cli.StringFlag{
			Name:    "dburl",
			Usage:   "url for database. e.g. postgres://user:password@localhost:5432/dbname",
			EnvVars: []string{"ENVD_DB_URL"},
		},
		&cli.BoolFlag{
			Name:    "no-auth",
			Usage:   "disable authentication. This is for development only. ",
			EnvVars: []string{"ENVD_NO_AUTH"},
			Aliases: []string{"n"},
		},
		&cli.StringFlag{
			Name:    "jwt-secret",
			Usage:   "secret for jwt token",
			Value:   "envd-server",
			EnvVars: []string{"ENVD_JWT_SECRET"},
			Aliases: []string{"js"},
		},
		&cli.DurationFlag{
			Name:    "jwt-expiration-timeout",
			Usage:   "expiration timeout for the issued jwt token",
			Value:   time.Hour * 24 * 365,
			EnvVars: []string{"ENVD_JWT_EXPIRATION_TIMEOUT"},
			Aliases: []string{"jet"},
		},
		&cli.StringFlag{
			Name:    "image-pull-secret-name",
			Usage:   "name of the image pull secret in the cluster",
			EnvVars: []string{"ENVD_IMAGE_PULL_SECRET_NAME"},
			Aliases: []string{"ipsn"},
		},
		&cli.BoolFlag{
			Name:    "resource-quota-enabled",
			Usage:   "enable resource quota",
			EnvVars: []string{"ENVD_RESOURCE_QUOTA_ENABLED"},
			Aliases: []string{"rqe"},
		},
	}
	internalApp.Action = runServer

	// Deal with debug flag.
	var debugEnabled bool

	internalApp.Before = func(context *cli.Context) error {
		debugEnabled = context.Bool("debug")

		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		if debugEnabled {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}
	return EnvdServerApp{
		App: internalApp,
	}
}

func runServer(clicontext *cli.Context) error {
	opt := server.Opt{
		Debug:                clicontext.Bool("debug"),
		KubeConfig:           clicontext.Path("kubeconfig"),
		HostKeyPath:          clicontext.Path("hostkey"),
		DBURL:                clicontext.String("dburl"),
		NoAuth:               clicontext.Bool("no-auth"),
		JWTSecret:            clicontext.String("jwt-secret"),
		ImagePullSecretName:  clicontext.String("image-pull-secret-name"),
		ResourceQuotaEnabled: clicontext.Bool("resource-quota-enabled"),
	}

	logrus.Debug("Starting server with options: ", opt)
	s, err := server.New(opt)
	if err != nil {
		return err
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil

}
