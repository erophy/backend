package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/golang-jwt/jwt"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/save", save)
	e.Logger.Fatal(e.Start(":1323"))
}

func save(c echo.Context) error {
	// Get email and password
	email := c.FormValue("email")
	password := c.FormValue("password")
	return c.String(http.StatusOK, "email:"+email+", password:"+password)
}

type TokenClaims struct {
	expiresAt int64
	token string
	claims string
	Issuer string
	jwt.StandardClaims

}

type JWTService interface {
	GenerateAccessToken(payload TokenClaims, jwtSecret []byte) (string, error)
}

type jwtService struct{}

func (s jwtService) GenerateAccessToken(payload TokenClaims, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	expiresAt:=time.Now().UTC().Add(60*time.Hour).Unix()
StandardClaims: jwt.StandardClaims{
	ExpiresAt: expiresAt,
	Issuer := <email>,
}
}
signedToken, err := token.SignedString(jwtSecret){
if err != nil {
return "", err
}

return signedToken, nil
}

func NewJWTService() JWTService {
	return &jwtService{}
}


