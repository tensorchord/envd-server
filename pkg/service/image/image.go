// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package image

import (
	"context"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/query"
)

type Service interface {
	CreateImageIfNotExist(ctx context.Context,
		owner, name string) (*types.ImageMeta, error)
	ListImages(ctx context.Context,
		owner string) ([]types.ImageMeta, error)
	GetImageByName(ctx context.Context,
		owner, name string) (*types.ImageMeta, error)
	GetImageByDigest(ctx context.Context,
		owner, digest string) (*types.ImageMeta, error)
}

type generalService struct {
	querier query.Querier
	logger  *logrus.Entry
}

func NewService(querier query.Querier) Service {
	return &generalService{
		querier: querier,
		logger:  logrus.WithField("service", "image"),
	}
}

func (g generalService) CreateImageIfNotExist(ctx context.Context,
	owner, name string) (*types.ImageMeta, error) {
	originalMeta, err := g.fetchMetadata(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "fetch metadata")
	}

	meta, err := g.GetImageByDigest(ctx, owner, originalMeta.Digest)
	if err != nil {
		if errdefs.IsNotFound(err) {
			// Create a new image
			if err := g.createImage(ctx, owner, &originalMeta); err != nil {
				return nil, errors.Wrap(err, "g.createImage")
			}
			return &originalMeta, nil
		} else {
			return nil, errors.Wrap(err, "g.GetImage")
		}
	}
	return meta, nil
}

func (g generalService) createImage(ctx context.Context,
	owner string, meta *types.ImageMeta) error {
	if meta == nil {
		return errors.New("meta is nil")
	}
	var pglabel, pgpackages, pgservices, pgpythonPackages pgtype.JSONB

	if err := pglabel.Set(meta.Labels); err != nil {
		return errors.Wrap(err, "pglabel.set")
	}

	if err := pgpackages.Set(meta.APTPackages); err != nil {
		return errors.Wrap(err, "pgpackages.set")
	}

	if err := pgservices.Set(meta.Ports); err != nil {
		return errors.Wrap(err, "pgservices.set")
	}

	if err := pgpythonPackages.Set(meta.PythonCommands); err != nil {
		return errors.Wrap(err, "pgpythonPackages.set")
	}

	_, err := g.querier.CreateImageInfo(ctx,
		query.CreateImageInfoParams{
			LoginName:    owner,
			Name:         meta.Name,
			Digest:       meta.Digest,
			Created:      meta.Created,
			Size:         meta.Size,
			Labels:       pglabel,
			AptPackages:  pgpackages,
			PypiCommands: pgpythonPackages,
			Services:     pgservices,
		})
	return err
}

func (g generalService) ListImages(ctx context.Context,
	owner string) ([]types.ImageMeta, error) {
	images, err := g.querier.ListImageByOwner(ctx, owner)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No image found
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "failed to list image by owner")
		}
	}
	res := []types.ImageMeta{}
	for _, info := range images {
		item, err := daoToImageMeta(info)
		if err != nil {
			return nil, errors.Wrap(err, "daoToImageMeta")
		}
		res = append(res, *item)
	}
	return res, nil
}

func (g generalService) GetImageByName(ctx context.Context,
	owner, name string) (*types.ImageMeta, error) {
	name, err := url.PathUnescape(name)
	if err != nil {
		return nil, errors.Wrap(err, "url.PathUnescape")
	}

	g.logger.WithFields(logrus.Fields{
		"owner": owner,
		"name":  name,
	}).Debug("getting the image by name")
	imageInfo, err := g.querier.GetImageInfoByName(ctx,
		query.GetImageInfoByNameParams{LoginName: owner, Name: name})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdefs.NotFound(err)
		} else {
			return nil, errors.Wrap(err, "g.querier.GetImageInfoByName")
		}
	}
	meta, err := daoToImageMeta(imageInfo)
	if err != nil {
		return nil, errors.Wrap(err, "util.DaoToImageMeta")
	}

	return meta, nil
}

func (g generalService) GetImageByDigest(ctx context.Context,
	owner, digest string) (*types.ImageMeta, error) {
	imageInfo, err := g.querier.GetImageInfoByDigest(ctx,
		query.GetImageInfoByDigestParams{LoginName: owner, Digest: digest})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdefs.NotFound(err)
		} else {
			return nil, errors.Wrap(err, "g.querier.GetImageByDigest")
		}
	}
	meta, err := daoToImageMeta(imageInfo)
	if err != nil {
		return nil, errors.Wrap(err, "util.DaoToImageMeta")
	}

	return meta, nil
}
