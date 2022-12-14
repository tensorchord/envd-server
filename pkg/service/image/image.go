package image

import (
	"context"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"

	"github.com/tensorchord/envd-server/api/types"
	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/query"
)

type Service interface {
	CreateImageIfNotExist(ctx context.Context,
		owner, name string) (*types.ImageMeta, error)
	ListImages(ctx context.Context,
		owner string) ([]types.ImageMeta, error)
	GetImage(ctx context.Context,
		owner, name string) (*types.ImageMeta, error)
}

type generalService struct {
	querier query.Querier
}

type CreateOptions struct {
	OwnerToken string `json:"owner_token"`
	Name       string `json:"name"`
	Digest     string `json:"digest"`
	Created    int64  `json:"created"`
	Size       int64  `json:"size"`
}

func NewService(querier query.Querier) Service {
	return &generalService{querier: querier}
}

func (g generalService) CreateImageIfNotExist(ctx context.Context,
	owner, name string) (*types.ImageMeta, error) {
	meta, err := g.GetImage(ctx, owner, name)
	if err != nil {
		if errdefs.IsNotFound(err) {
			// Create a new image
			meta, err = g.createImage(ctx, owner, name)
			if err != nil {
				return nil, errors.Wrap(err, "g.createImage")
			}
			return meta, nil
		} else {
			return nil, errors.Wrap(err, "g.GetImage")
		}
	}
	return meta, nil
}

func (g generalService) createImage(ctx context.Context,
	owner, name string) (*types.ImageMeta, error) {
	meta, err := g.fetchMetadata(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "fetch metadata")
	}
	var pglabel pgtype.JSONB
	err = pglabel.Set(meta.Labels)
	if err != nil {
		return nil, errors.Wrap(err, "pglabel.set")
	}

	_, err = g.querier.CreateImageInfo(context.Background(),
		query.CreateImageInfoParams{OwnerToken: owner,
			Name: meta.Name, Digest: meta.Digest,
			Created: meta.Created, Size: meta.Size, Labels: pglabel})
	return &meta, err
}

func (g generalService) ListImages(ctx context.Context,
	owner string) ([]types.ImageMeta, error) {
	images, err := g.querier.ListImageByOwner(context.Background(), owner)
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

func (g generalService) GetImage(ctx context.Context,
	owner, name string) (*types.ImageMeta, error) {
	name, err := url.PathUnescape(name)
	if err != nil {
		return nil, errors.Wrap(err, "url.PathUnescape")
	}

	imageInfo, err := g.querier.GetImageInfo(context.Background(), query.GetImageInfoParams{OwnerToken: owner, Name: name})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdefs.NotFound(err)
		} else {
			return nil, errors.Wrap(err, "g.querier.GetImageInfo")
		}
	}
	meta, err := daoToImageMeta(imageInfo)
	if err != nil {
		return nil, errors.Wrap(err, "util.DaoToImageMeta")
	}

	return meta, nil
}
