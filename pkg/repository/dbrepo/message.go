package dbrepo

import (
	"context"

	"github.com/chiachun0920/platform-api/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageDBRepo struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

func NewMessageDBRepo(db *mongo.Client, dbName string) *MessageDBRepo {
	return &MessageDBRepo{db: db, dbName: dbName, collectionName: "message"}
}

func (repo *MessageDBRepo) SaveMessage(m *dto.Message) error {
	coll := repo.db.Database(repo.dbName).Collection(repo.collectionName)

	_, err := coll.InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MessageDBRepo) ListMessages(
	customerId string,
) ([]*dto.Message, error) {
	coll := repo.db.Database(repo.dbName).Collection(repo.collectionName)
	cursor, err := coll.Find(
		context.TODO(),
		bson.D{{Key: "sender", Value: customerId}},
	)
	if err != nil {
		return nil, err
	}

	var messages []*dto.Message
	if err = cursor.All(context.TODO(), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
