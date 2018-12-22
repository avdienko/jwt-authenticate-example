package main

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

var secret = []byte("secret")

func createToken() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &User{
		Name: "Dima",
		Age:  26,
	})

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		log.Fatalln(err)
	}

	return tokenStr
}

func main() {

	tokenstring := createToken()
	// This is token
	log.Println(tokenstring)

	// Parse this by the secrete
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// When using `Parse`, the result a map
	log.Println(token.Claims, err)

	//  Decode token to struct
	user := User{}
	token, err = jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	log.Println(token.Valid, user, err)
}
