package utils

import (
	"fmt"
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

//Verify Token
func VerifyToken(tokenString string) (userID int64, err error) {
	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
		fmt.Println("----> tokenString:", tokenString)
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return -1, err
	}

	tokenIsValid := parsedToken.Valid


	if !tokenIsValid {
		return -1, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, err
	}

	// email := claims["email"].(string)
	userId := int64(claims["user_id"].(float64))

	return userId, nil
}