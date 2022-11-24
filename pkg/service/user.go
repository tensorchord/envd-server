// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/tensorchord/envd-server/pkg/query"
)

type UserService struct {
	querier query.Querier
}

func NewUserService(querier query.Querier) *UserService {
	return &UserService{
		querier: querier,
	}
}

func (userService *UserService) Register(IdentityToken string, PublicKey []byte) error {
	_, err := userService.querier.CreateUser(context.Background(), query.CreateUserParams{IdentityToken: IdentityToken, PublicKey: PublicKey})
	return err
}

func (userService *UserService) GetPubKey(IdentityToken string) ([]byte, error) {
	user, err := userService.querier.GetUser(context.Background(), IdentityToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return user.PublicKey, nil
}

func (userService *UserService) Auth(IdentityToken string) (bool, error) {
	_, err := userService.querier.GetUser(context.Background(), IdentityToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
