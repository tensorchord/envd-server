// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package app

import (
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
	s, err := server.New(server.Opt{
		Debug:       clicontext.Bool("debug"),
		KubeConfig:  clicontext.Path("kubeconfig"),
		HostKeyPath: clicontext.Path("hostkey"),
	})
	if err != nil {
		return err
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil

}
