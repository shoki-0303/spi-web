package helpers

import (
	"fmt"
	"spi-web/app/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJWTtoken(adminUser *models.AdminUser) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = adminUser.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	jwt, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Printf("action=CreateJWTtoken err=%s", err)
	}
	return jwt
}
