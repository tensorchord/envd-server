// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package user

import "golang.org/x/crypto/bcrypt"

func GenerateHashedSaltPassword(pwd []byte) ([]byte, error) {
	bcryptedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return bcryptedPwd, nil
}

func CompareHashAndPassword(hashedPwd, pwd []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPwd, pwd)
}
