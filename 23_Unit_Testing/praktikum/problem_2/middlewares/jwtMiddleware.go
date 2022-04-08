package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/constants"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	/*
		Bila ingin membatas kadaluarsa
		 claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	*/

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
