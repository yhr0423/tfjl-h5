package db

import (
	"context"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (manager *dbManager) SetActivityCollection(collection string) {
	manager.ActivityCollection = manager.TFJLDatabase.Collection(collection)
	logrus.Info("Set Collection:ActivityCollection success!")
}

func (manager *dbManager) FindActivitys(filter bson.M, result *[]protocols.T_Activity_Data, opts ...*options.FindOptions) error {
	cursor, err := manager.ActivityCollection.Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	// 通过All一次性获取所有结果
	if err = cursor.All(context.Background(), result); err != nil {
		return err
	}

	return nil
}
