package auth

import (
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var GetTokenHandler = http.HandlerFunc(getTokenHandleFunc)

func getTokenHandleFunc(w http.ResponseWriter, r *http.Request) {
	// set header
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = true
	claims["name"] = true
	claims["iat"] = true
	claims["exp"] = true

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SIGNINGKEY")))

	// JWTを返却
	w.Write([]byte(tokenString))
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
