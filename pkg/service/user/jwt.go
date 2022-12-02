package user

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTIssuer struct {
	ExpirationTime time.Duration
	Key            string
}

func newJWTIssuer(expirationTime time.Duration, key string) *JWTIssuer {
	return &JWTIssuer{
		ExpirationTime: expirationTime,
		Key:            key,
	}
}

func (j *JWTIssuer) Issue(username string) (string, error) {
	// Create the Claims
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Key))
}

func (j *JWTIssuer) Validate(token string) (string, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Key), nil
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to parse token")
	}
	if !tkn.Valid {
		return "", fmt.Errorf("failed to parse the claims")
	}

	return claims.Username, nil
}
