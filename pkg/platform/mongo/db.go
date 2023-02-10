package models

import (
	"context"
	. "my-app/pkg/config"
	. "my-app/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Close(client *mongo.Client, ctx context.Context) {
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, err
}

func Insert(config Configuration, user interface{}) (*mongo.InsertOneResult, error) {
	currDateTime := time.Now()
	doc := bson.D{
		{Key: "configKey", Value: config.Key},
		{Key: "configValue", Value: config.Value},
		{Key: "createdBy", Value: user.(string)},
		{Key: "creationTime", Value: currDateTime},
		{Key: "modifiedBy", Value: user.(string)},
		{Key: "modificationTime", Value: currDateTime},
	}
	collection := Client.Database(Db).Collection(Col)
	result, err := collection.InsertOne(Ctx, doc)
	return result, err
}

func Query(configKey string) (ConfigurationData, error) {
	var config ConfigurationData
	query := bson.D{
		{Key: "configKey", Value: configKey},
	}
	field := bson.D{{Key: "_id", Value: 0}}

	collection := Client.Database(Db).Collection(Col)
	res := collection.FindOne(Ctx, query, options.FindOne().SetProjection(field))
	err := res.Decode(&config)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return config, nil
		}
	}

	return config, err
}

func Delete(configKey string) (result *mongo.DeleteResult, err error) {
	query := bson.D{
		{Key: "configKey", Value: configKey},
	}
	collection := Client.Database(Db).Collection(Col)
	result, err = collection.DeleteOne(Ctx, query)
	return result, err
}

func Update(config Configuration, user interface{}) (result *mongo.UpdateResult, err error) {
	filter := bson.D{
		{Key: "configKey", Value: config.Key},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "configValue", Value: config.Value},
			{Key: "modifiedBy", Value: user.(string)},
			{Key: "modificationTime", Value: time.Now()},
		}},
	}
	collection := Client.Database(Db).Collection(Col)
	result, err = collection.UpdateOne(Ctx, filter, update)
	return result, err
}
