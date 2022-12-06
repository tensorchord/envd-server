// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server

const (
	ContextLoginName = "login-name"
)

type AuthMiddlewareHeaderRequest struct {
	JWTToken string `header:"Authorization"`
}

type AuthMiddlewareURIRequest struct {
	LoginName string `uri:"login_name" example:"alice"`
}
