package dbrepo

import (
	"context"

	"github.com/chiachun0920/platform-api/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerDBRepo struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

func NewCustomerDBRepo(db *mongo.Client, dbName string) *CustomerDBRepo {
	return &CustomerDBRepo{db: db, dbName: dbName, collectionName: "customer"}
}

func (repo *CustomerDBRepo) UpsertCustomerProfile(
	profile *dto.CustomerProfile,
) error {
	profileByte, err := bson.Marshal(*profile)

	if err != nil {
		return err
	}

	var update bson.M
	err = bson.Unmarshal(profileByte, &update)
	if err != nil {
		return err
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "userId", Value: profile.UserID}}

	coll := repo.db.Database(repo.dbName).Collection("user")

	_, err = coll.UpdateOne(
		context.TODO(),
		filter,
		bson.D{{Key: "$set", Value: update}},
		opts,
	)
	if err != nil {
		return err
	}

	return nil
}
