package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var JwtKey = []byte("my_secret_key")

var Users = make(map[string]string)

var Client *mongo.Client
var Ctx context.Context

var Db, Col string
