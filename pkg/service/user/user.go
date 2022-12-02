// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package user

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/tensorchord/envd-server/pkg/query"
)

var (
	// ExpirationTimeDefault is the default expiration time of the token.
	expirationTimeDefault = time.Hour * 24 * 365
	// jwtSecretDefault is the secret used to sign the token.
	jwtSecretDefault = "envd-secret"
)

type Service struct {
	querier   query.Querier
	jwtIssuer *JWTIssuer
}

func NewService(querier query.Querier) *Service {
	return &Service{
		querier:   querier,
		jwtIssuer: newJWTIssuer(expirationTimeDefault, jwtSecretDefault),
	}
}

func (u *Service) Register(loginName string,
	pwd, PublicKey []byte) (string, error) {
	bcryptedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	_, err = u.querier.CreateUser(
		context.Background(), query.CreateUserParams{
			LoginName:    loginName,
			PasswordHash: string(bcryptedPwd),
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

func (u *Service) Login(loginName string, pwd []byte) (bool, string, error) {
	rawUser, err := u.querier.GetUser(context.Background(), loginName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, "", nil
		} else {
			return false, "", err
		}
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(rawUser.PasswordHash), pwd); err != nil {
		return false, "", err
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
