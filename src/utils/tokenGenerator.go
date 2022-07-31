/*
USAGE:
1. Install the "github.com/golang-jwt/jwt" package
2. Create a new envvar `JWT_SECRET`
3. Create a `TokenClaims` instance and fill it. Example:

    expiresAt := <google how to get the minutes in unix. Expiration time you could choose by yourself>.
    claims := TokenClaims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expiresAt,
            Issuer:    <email>,
        }
    }

4. Pass the claims and the jwt secret (google how to read the envvar as well) into the function call
*/

package utils

import (
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
}

type JWTService interface {
	GenerateAccessToken(payload TokenClaims, jwtSecret []byte) (string, error)
}

type jwtService struct{}

func (s jwtService) GenerateAccessToken(payload TokenClaims, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func NewJWTService() JWTService {
	return &jwtService{}
}
