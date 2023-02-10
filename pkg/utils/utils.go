package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	. "my-app/pkg/config"
	. "my-app/pkg/models"

	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func AuthenticateUser(req *http.Request) (interface{}, bool) {
	reqToken := req.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Println(errors.New("Unauthorized:Bad authorization header!!"))
		return nil, false
	}
	reqToken = splitToken[1]

	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return nil, false
	}

	return claims["username"], true
}

func ValidateKey(key string) bool {
	if key == "" {
		return false
	}
	return true
}

func SendErrorResponse(apiRes ApiResponse, errCode int, rw http.ResponseWriter) {
	log.Println(apiRes.Msg)
	rw.WriteHeader(http.StatusBadRequest)
	rw.Header().Set("Content-Type", "application/json")
	jsonRes, err := json.Marshal(apiRes)
	if err != nil {
		log.Println(err.Error())
		io.WriteString(rw, "Something went wrong!!")
		return
	}
	rw.Write(jsonRes)
}

func SendResponse(apiRes ApiResponse, rw http.ResponseWriter) {
	jsonRes, err := json.Marshal(apiRes)
	if err != nil {
		log.Println(err.Error())
		io.WriteString(rw, "Something went wrong!!")
		return
	}
	rw.Write(jsonRes)
}
