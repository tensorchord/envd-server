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
