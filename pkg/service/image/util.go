// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package image

import (
	"encoding/json"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/query"
)

func daoToImageMeta(dao query.ImageInfo) (*types.ImageMeta, error) {
	var label map[string]string
	if err := dao.Labels.AssignTo(&label); err != nil {
		return nil, err
	}

	var aptPackages []string
	if err := dao.AptPackages.AssignTo(&aptPackages); err != nil {
		return nil, err
	}

	var pythonCommands []string
	if err := dao.PypiCommands.AssignTo(&pythonCommands); err != nil {
		return nil, err
	}

	var ports []types.EnvironmentPort
	if err := dao.Services.AssignTo(&ports); err != nil {
		return nil, err
	}

	meta := types.ImageMeta{
		Name:           dao.Name,
		Digest:         dao.Digest,
		Created:        dao.Created,
		Size:           dao.Size,
		Labels:         label,
		APTPackages:    aptPackages,
		PythonCommands: pythonCommands,
		Ports:          ports,
	}
	return &meta, nil
}

func portsFromLabel(label string) ([]types.EnvironmentPort, error) {
	var ports []types.EnvironmentPort
	if err := json.Unmarshal([]byte(label), &ports); err != nil {
		return nil, err
	}

	return ports, nil
}

func aptPackagesFromLabel(label string) ([]string, error) {
	var packages []string
	if err := json.Unmarshal([]byte(label), &packages); err != nil {
		return nil, err
	}
	return packages, nil
}

func pythonCommandsFromLabel(label string) ([]string, error) {
	var commands []string
	if err := json.Unmarshal([]byte(label), &commands); err != nil {
		return nil, err
	}
	return commands, nil
}
