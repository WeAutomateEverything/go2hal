package telegram

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"
)

func makeToken(roomid uint32) (string, error) {
	claims := CustomClaims{
		roomid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 120).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func Escape(request string) string {
	request = strings.Replace(request, "*", "\\*", -1)
	request = strings.Replace(request, "_", "\\_", -1)
	return request
}
