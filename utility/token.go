package utility

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Secrete_Key = []byte("MRVNfviwapQJjczZhCXceLfntyxxeB")

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(username string) (string, error) {

	expitation := time.Now().Add(time.Minute * 15)

	clams := Claims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expitation.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)

	tokenString, err := token.SignedString(Secrete_Key)
	if err != nil {
		fmt.Println("Erro occured ", err)
		return "", err
	}

	return tokenString, nil
}
