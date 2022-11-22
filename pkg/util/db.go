// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package util

import (
	"context"
	_ "embed"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

//go:embed sql/schema.sql
var schemaSql string

func ApplySchema(conn *pgx.Conn) error {
	tag, err := conn.Exec(context.Background(), schemaSql)
	logrus.Debugf("apply schema result: %s", tag)
	return err
}
