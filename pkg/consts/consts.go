// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package consts

const (
	EnvdLabelPrefix = "ai.tensorchord.envd."

	PodLabelUID               = EnvdLabelPrefix + "uid"
	PodLabelEnvironmentName   = EnvdLabelPrefix + "environment-name"
	PodLabelJupyterAddr       = EnvdLabelPrefix + "jupyter.address"
	PodLabelRStudioServerAddr = EnvdLabelPrefix + "rstudio.server.address"

	ImageLabelContainerName  = EnvdLabelPrefix + "container.name"
	ImageLabelPorts          = EnvdLabelPrefix + "ports"
	ImageLabelRepo           = EnvdLabelPrefix + "repo"
	ImageLabelAPTPackages    = EnvdLabelPrefix + "apt.packages"
	ImageLabelPythonCommands = EnvdLabelPrefix + "pypi.commands"
)
