// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/tensorchord/envd-server/api/types"
)

var (
	ErrNoAuth = errors.New("no authentication provided")
	ErrNoUser = errors.New("no user provided")
)

// Register authenticates the envd server.
// It returns unauthorizedError when the authentication fails.
func (cli *Client) Register(ctx context.Context, auth types.AuthNRequest) (types.AuthNResponse, error) {
	resp, err := cli.post(ctx, "/register", url.Values{}, auth, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.AuthNResponse{}, err
	}

	var response types.AuthNResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}

// Login logins the envd server.
// It returns unauthorizedError when the authentication fails.
func (cli *Client) Login(ctx context.Context, auth types.AuthNRequest) (types.AuthNResponse, error) {
	resp, err := cli.post(ctx, "/login", url.Values{}, auth, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return types.AuthNResponse{}, err
	}

	var response types.AuthNResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}

func (cli Client) getUserAndHeaders() (string, map[string][]string, error) {
	if !cli.auth {
		if cli.user == "" {
			return "", nil, ErrNoUser
		}
		return cli.user, nil, nil
	}

	if cli.jwtToken == "" || cli.user == "" {
		return "", nil, ErrNoAuth
	}
	return cli.user, map[string][]string{
		"Authorization": {cli.jwtToken},
	}, nil
}
