package token

import "github.com/golang-jwt/jwt/v4"

var (
	secretKey = []byte("nightsea-9527")
	hashkey   = jwt.SigningMethodHS512
)

type cliams struct {
	Data   interface{} `json:"data"`
	Cliams *jwt.StandardClaims
}
