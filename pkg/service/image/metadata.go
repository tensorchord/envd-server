// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package image

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/image"
	containertypes "github.com/containers/image/v5/types"
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/pkg/consts"
)

func (g generalService) fetchMetadata(ctx context.Context, imageName string) (
	meta types.ImageMeta, err error) {
	ref, err := docker.ParseReference(fmt.Sprintf("//%s", imageName))
	if err != nil {
		return meta, errors.Wrap(err, "docker.ParseReference")
	}
	sys := &containertypes.SystemContext{}
	src, err := ref.NewImageSource(ctx, sys)
	if err != nil {
		return meta, errors.Wrap(err, "ref.NewImageSource")
	}
	defer src.Close()
	digest, err := docker.GetDigest(ctx, sys, ref)
	if err != nil {
		return meta, errors.Wrap(err, "docker.GetDigest")
	}
	image, err := image.FromUnparsedImage(ctx, sys, image.UnparsedInstance(src, &digest))
	if err != nil {
		return meta, errors.Wrap(err, "image.FromUnparsedImage")
	}
	inspect, err := image.Inspect(ctx)
	if err != nil {
		return meta, errors.Wrap(err, "image.Inspect")
	}

	// correct the image size
	var size int64
	for _, layer := range inspect.LayersData {
		size += layer.Size
	}

	meta = types.ImageMeta{
		Name:    imageName,
		Created: inspect.Created.Unix(),
		Digest:  string(digest),
		Labels:  inspect.Labels,
		Size:    size,
	}

	// Get the apt packages
	if aptLabel, ok := inspect.Labels[consts.ImageLabelAPTPackages]; ok {
		packages, err := aptPackagesFromLabel(aptLabel)
		if err != nil {
			return meta, errors.Wrap(err, "apt packages")
		}
		meta.APTPackages = packages
	}

	// Get the pip packages
	if pipLabel, ok := inspect.Labels[consts.ImageLabelPythonCommands]; ok {
		commands, err := pythonCommandsFromLabel(pipLabel)
		if err != nil {
			return meta, errors.Wrap(err, "pip packages")
		}
		meta.PythonCommands = commands
	}

	// Get the services
	if servicesLabel, ok := inspect.Labels[consts.ImageLabelPorts]; ok {
		ports, err := portsFromLabel(servicesLabel)
		if err != nil {
			return meta, errors.Wrap(err, "services")
		}
		meta.Ports = ports
	}

	logrus.WithField("image meta", meta).Debug("get image meta before creating env")
	return meta, nil
}
