package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Configuration struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type ConfigurationData struct {
	Key              string             `bson:"configKey"`
	Value            interface{}        `bson:"configValue"`
	CreatedBy        string             `bson:"createdBy"`
	CreationTime     primitive.DateTime `bson:"creationTime"`
	ModifiedBy       string             `bson:"modifiedBy"`
	ModificationTime primitive.DateTime `bson:"modificationTime"`
	// add project name and make key+project as unique key
}

type ApiResponse struct {
	Data interface{}
	Msg  string
}
