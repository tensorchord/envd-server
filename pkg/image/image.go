package image

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/image"
	containertypes "github.com/containers/image/v5/types"
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/api/types"
)

// TODO(gaocegege): Support image registry auth.
func FetchMetadata(ctx context.Context, imageName string) (
	meta types.ImageMeta, err error) {
	ref, err := docker.ParseReference(fmt.Sprintf("//%s", imageName))
	if err != nil {
		return
	}
	sys := &containertypes.SystemContext{}
	src, err := ref.NewImageSource(ctx, sys)
	if err != nil {
		return
	}
	digest, err := docker.GetDigest(ctx, sys, ref)
	if err != nil {
		return
	}
	image, err := image.FromUnparsedImage(ctx, sys, image.UnparsedInstance(src, &digest))
	if err != nil {
		return
	}
	inspect, err := image.Inspect(ctx)
	if err != nil {
		return
	}

	// correct the image size
	var size int64 = 0
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
	logrus.WithField("image meta", meta).Debug("get image meta before creating env")
	src.Close()
	return meta, nil
}
