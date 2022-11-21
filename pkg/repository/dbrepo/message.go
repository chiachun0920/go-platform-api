package dbrepo

import (
	"context"

	"github.com/chiachun0920/platform-api/pkg/dto"
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

func (repo *MessageDBRepo) ListMessages() ([]*dto.Message, error) {
	return nil, nil
}
