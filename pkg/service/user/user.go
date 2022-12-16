// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package user

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	"github.com/tensorchord/envd-server/errdefs"
	"github.com/tensorchord/envd-server/pkg/query"
)

type Service interface {
	Register(loginName, pwd string,
		publicKey []byte) (string, error)
	GetPubKey(loginName string) ([]byte, error)
	Login(loginName, pwd string, auth bool) (bool, string, error)

	ValidateJWT(token string) (string, error)
}

type generalService struct {
	querier   query.Querier
	jwtIssuer *JWTIssuer
}

func NewService(querier query.Querier,
	secret string, expirationTimeDefault time.Duration) Service {
	return &generalService{
		querier:   querier,
		jwtIssuer: newJWTIssuer(expirationTimeDefault, secret),
	}
}

func (u *generalService) Register(loginName, pwd string,
	publicKey []byte) (string, error) {
	hashed, err := GenerateHashedSaltPassword([]byte(pwd))
	if err != nil {
		return "", err
	}
	_, err = u.querier.CreateUser(
		context.Background(), query.CreateUserParams{
			LoginName:    loginName,
			PasswordHash: string(hashed),
			PublicKey:    publicKey,
		})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return "", errdefs.Conflict(errors.New("login name already exists"))
			}
		}
		return "", err
	}

	token, err := u.jwtIssuer.Issue(loginName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *generalService) GetPubKey(loginName string) ([]byte, error) {
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

func (u *generalService) Login(loginName, pwd string, auth bool) (bool, string, error) {
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

func (u *generalService) ValidateJWT(token string) (string, error) {
	return u.jwtIssuer.Validate(token)
}
