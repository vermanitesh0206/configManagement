package controllers

import (
	"encoding/json"
	. "my-app/pkg/config"
	. "my-app/pkg/models"
	. "my-app/pkg/platform/mongo"
	. "my-app/pkg/utils"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func Login(rw http.ResponseWriter, req *http.Request) {
	var apiRes ApiResponse
	var creds Credentials
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusBadRequest, rw)
		return
	}

	expectedPassword, ok := Users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		apiRes.Msg = "Username/Password incorrect!!"
		SendErrorResponse(apiRes, http.StatusUnauthorized, rw)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusInternalServerError, rw)
		return
	}

	apiRes.Data = tokenString
	apiRes.Msg = "Token generated successfully!!"
	SendResponse(apiRes, rw)
}

func AddConfig(rw http.ResponseWriter, req *http.Request) {
	var apiRes ApiResponse
	var user interface{}
	var authRes bool
	if user, authRes = AuthenticateUser(req); !authRes {
		apiRes.Msg = "Unauthorized:please login first!!"
		SendErrorResponse(apiRes, http.StatusUnauthorized, rw)
		return
	}

	var config Configuration
	err := json.NewDecoder(req.Body).Decode(&config)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusBadRequest, rw)
		return
	}

	if validKey := ValidateKey(config.Key); !validKey {
		apiRes.Msg = "Key can't be empty!!"
		SendErrorResponse(apiRes, http.StatusBadRequest, rw)
		return
	}

	_, err = Insert(config, user)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusInternalServerError, rw)
		return
	}

	apiRes.Data = config
	apiRes.Msg = "Configuration inserted successfully!!"
	SendResponse(apiRes, rw)
}

func UpdateConfig(rw http.ResponseWriter, req *http.Request) {
	var apiRes ApiResponse
	var user interface{}
	var authRes bool
	if user, authRes = AuthenticateUser(req); !authRes {
		apiRes.Msg = "Unauthorized:please login first!!"
		SendErrorResponse(apiRes, http.StatusUnauthorized, rw)
		return
	}

	var config Configuration
	err := json.NewDecoder(req.Body).Decode(&config)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusBadRequest, rw)
		return
	}

	if validKey := ValidateKey(config.Key); !validKey {
		apiRes.Msg = "Key can't be empty!!"
		SendErrorResponse(apiRes, http.StatusBadRequest, rw)
		return
	}

	result, err := Update(config, user)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusInternalServerError, rw)
		return
	}

	apiRes.Data = config
	if result.MatchedCount == 0 {
		apiRes.Msg = "Configuration not found!!"
	} else {
		apiRes.Msg = "Configuration updated successfully!!"
	}
	SendResponse(apiRes, rw)
}

func SearchConfig(rw http.ResponseWriter, req *http.Request) {
	var apiRes ApiResponse
	if _, authRes := AuthenticateUser(req); !authRes {
		apiRes.Msg = "Unauthorized:please login first!!"
		SendErrorResponse(apiRes, http.StatusUnauthorized, rw)
		return
	}

	vars := mux.Vars(req)
	configKey := vars["configKey"]

	config, err := Query(configKey)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusInternalServerError, rw)
		return
	}

	if config.Value == nil {
		apiRes.Data = nil
		apiRes.Msg = "Configuration not found!!"
	} else {
		apiRes.Data = config
		apiRes.Msg = "Configuration found!!"
	}
	SendResponse(apiRes, rw)
}

func DeleteConfig(rw http.ResponseWriter, req *http.Request) {
	var apiRes ApiResponse
	if _, authRes := AuthenticateUser(req); !authRes {
		apiRes.Msg = "Unauthorized:please login first!!"
		SendErrorResponse(apiRes, http.StatusUnauthorized, rw)
		return
	}

	vars := mux.Vars(req)
	configKey := vars["configKey"]

	result, err := Delete(configKey)
	if err != nil {
		apiRes.Msg = err.Error()
		SendErrorResponse(apiRes, http.StatusInternalServerError, rw)
		return
	}

	apiRes.Data = configKey
	if result.DeletedCount == 0 {
		apiRes.Msg = "Configuration not found!!"
	} else {
		apiRes.Msg = "Configuration deleted successfully!!"
	}
	SendResponse(apiRes, rw)
}
