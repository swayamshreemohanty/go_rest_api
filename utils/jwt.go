package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey="dslfhdsiofhdiosfhioshdfioh"

func GenerateToken(email string, userId int64) (string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
	
}