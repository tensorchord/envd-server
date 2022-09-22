// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/tensorchord/envd-server/pkg/app"
	"github.com/tensorchord/envd-server/pkg/version"
)

func run(args []string) error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Name, version.Package, c.App.Version, version.Revision)
	}

	app := app.New()
	return app.Run(args)
}

func handleErr(err error) {
	if err == nil {
		return
	}

	logrus.Error(err)
}

func main() {
	err := run(os.Args)
	handleErr(err)
}
