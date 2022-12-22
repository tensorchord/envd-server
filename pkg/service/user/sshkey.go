// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package user

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgconn"
	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/query"
)

func (u *generalService) ListPubKeys(ctx context.Context, loginName string) ([]query.Key, error) {
	return u.querier.ListKeys(ctx, loginName)
}

func (u *generalService) CreatePubKey(ctx context.Context, loginName, name string, pubKey []byte) error {
	_, err := u.querier.CreateKey(ctx, query.CreateKeyParams{
		LoginName: loginName,
		PublicKey: pubKey,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return errdefs.Conflict(errors.New("key already exists"))
			}
		}
		return err
	}
	return err
}

func (u *generalService) GetPubKey(ctx context.Context, loginName, name string) ([]byte, error) {
	k, err := u.querier.GetKey(ctx, query.GetKeyParams{
		LoginName: loginName,
		Name:      name,
	})
	if err != nil {
		return nil, err
	}

	return k.PublicKey, nil
}
