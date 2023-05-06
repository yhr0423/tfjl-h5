package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) FindOne(collection string, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	return manager.TFJLDatabase.Collection(collection).FindOne(context.Background(), filter, opts...).Decode(result)
}

func (manager *dbManager) Find(collection string, filter interface{}, result interface{}, opts ...*options.FindOptions) error {
	cursor, err := manager.TFJLDatabase.Collection(collection).Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	if err = cursor.Err(); err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), result); err != nil {
		return err
	}
	return nil
}

func (manager *dbManager) Count(collection string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := manager.TFJLDatabase.Collection(collection).CountDocuments(context.Background(), filter, opts...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (manager *dbManager) UpdateOne(collection string, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return manager.TFJLDatabase.Collection(collection).UpdateOne(context.Background(), filter, update, opts...)
}

func (manager *dbManager) InsertOne(collection string, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return manager.TFJLDatabase.Collection(collection).InsertOne(context.Background(), data, opts...)
}

func (manager *dbManager) InsertMany(collection string, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return manager.TFJLDatabase.Collection(collection).InsertMany(context.Background(), data, opts...)
}
