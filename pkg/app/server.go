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
			Name:    "kube-config",
			Usage:   "kube-config path",
			EnvVars: []string{"KUBE_CONFIG", "KUBECONFIG"},
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
		Debug:      clicontext.Bool("debug"),
		KubeConfig: clicontext.Path("kube-config"),
	})
	if err != nil {
		return err
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil

}
