// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package user

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/tensorchord/envd-server/pkg/query"
)

type Service struct {
	querier   query.Querier
	jwtIssuer *JWTIssuer
}

func NewService(querier query.Querier,
	secret string, expirationTimeDefault time.Duration) *Service {
	return &Service{
		querier:   querier,
		jwtIssuer: newJWTIssuer(expirationTimeDefault, secret),
	}
}

func (u *Service) Register(loginName, pwd string,
	PublicKey []byte) (string, error) {
	hashed, err := GenerateHashedSaltPassword([]byte(pwd))
	if err != nil {
		return "", err
	}
	_, err = u.querier.CreateUser(
		context.Background(), query.CreateUserParams{
			LoginName:    loginName,
			PasswordHash: string(hashed),
			PublicKey:    PublicKey,
		})

	if err != nil {
		return "", err
	}

	token, err := u.jwtIssuer.Issue(loginName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *Service) GetPubKey(loginName string) ([]byte, error) {
	user, err := u.querier.GetUser(context.Background(), loginName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return user.PublicKey, nil
}

func (u *Service) Login(loginName, pwd string, auth bool) (bool, string, error) {
	if auth {
		rawUser, err := u.querier.GetUser(context.Background(), loginName)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return false, "", nil
			} else {
				return false, "", err
			}
		}

		if err := CompareHashAndPassword(
			[]byte(rawUser.PasswordHash), []byte(pwd)); err != nil {
			return false, "", err
		}
	}

	token, err := u.jwtIssuer.Issue(loginName)
	if err != nil {
		return false, "", err
	}
	return true, token, nil
}

func (u *Service) ValidateJWT(token string) (string, error) {
	return u.jwtIssuer.Validate(token)
}
