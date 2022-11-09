// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package consts

const (
	EnvdLabelPrefix = "envd.tensorchord.ai/"

	LabelUID             = EnvdLabelPrefix + "uid"
	LabelEnvironmentName = EnvdLabelPrefix + "environment-name"

	ImageLabelContainerName     = "ai.tensorchord.envd.container.name"
	ImageLabelPorts             = "ai.tensorchord.envd.ports"
	ImageLabelRepo              = "ai.tensorchord.envd.repo"
	ImageLabelJupyterAddr       = "ai.tensorchord.envd.jupyter.address"
	ImageLabelRStudioServerAddr = "ai.tensorchord.envd.rstudio.server.address"
)
