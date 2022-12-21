// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package query

import (
	"context"
)

type Querier interface {
	CreateImageInfo(ctx context.Context, arg CreateImageInfoParams) (ImageInfo, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteUser(ctx context.Context, id int64) error
	GetImageInfoByDigest(ctx context.Context, arg GetImageInfoByDigestParams) (ImageInfo, error)
	GetImageInfoByName(ctx context.Context, arg GetImageInfoByNameParams) (ImageInfo, error)
	GetUser(ctx context.Context, loginName string) (User, error)
	ListImageByOwner(ctx context.Context, loginName string) ([]ImageInfo, error)
	ListUsers(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)
