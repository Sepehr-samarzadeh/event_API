package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret" //just for the demo we need better way to make this key!

//generate the JWT token

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), //expire time for token (from now till added value) 2hours
	}) //generate new token with data attached to it
	return token.SignedString([]byte(secretKey)) //the  result is a complicated value so i turned it to normal string so we can deal with it
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // to check if a value stored on that field is the  type HMAC...
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return errors.New("Could not parse token!")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("Invalid token!")
	}

	//claims,ok := parsedToken.Claims.(jwt.MapClaims)

	//if !ok{
	//return errors.New("Invalid token claims!")
	//}

	//email:= claims["email"].(string) //type checking (check if email is type string)
	//userId:= claims["userId"].(int64)

	return nil
}
