package repository

import (
	"context"
	"fmt"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DataMongo struct {
	client *mongo.Client
	dbName string
}

func NewDataMongoWithConfig(config configs.ConfigDatabaseMongo) DataMongoManager {
	mongoClient, err := mongo.Connect(
		options.Client().
			ApplyURI(config.URI).
			SetAuth(options.Credential{
				Username: config.User,
				Password: config.Password,
			}))
	if err != nil {
		panic(fmt.Sprintf("error to connect in database: %v", err))
	}

	return &DataMongo{
		client: mongoClient,
		dbName: config.DbName,
	}
}

func (d *DataMongo) Begin() (TransactionMongoManager, error) {
	session, err := d.client.StartSession()
	if err != nil {
		return nil, fmt.Errorf("error starting session: %v", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}

	return newTransactionMongo(*session, d.dbName), nil
}

func (d *DataMongo) Close() error {
	if err := d.client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("error disconnecting from database: %v", err)
	}
	return nil
}

func (d *DataMongo) IsHealthy(ctx context.Context) (bool, error) {
	err := d.client.Ping(ctx, nil)
	if err != nil {
		return false, errors.Wrap(err)
	}
	return true, nil
}

func (d *DataMongo) Command() ConnectorMongo {
	return ConnectorMongo{
		Aggregate: func(ctx context.Context, collection string, pipeline any) (RowsMongo, error) {
			cursor, err := d.client.Database(d.dbName).Collection(collection).Aggregate(ctx, pipeline)
			if err != nil {
				return nil, errors.Wrap(err)
			}
			return cursor, nil
		},
		Find: func(ctx context.Context, collection string, filter any) (RowsMongo, error) {
			cursor, err := d.client.Database(d.dbName).Collection(collection).Find(ctx, filter)
			if err != nil {
				return nil, errors.Wrap(err)
			}
			return cursor, nil
		},
		FindOne: func(ctx context.Context, collection string, filter any, result any) error {
			err := d.client.Database(d.dbName).Collection(collection).FindOne(ctx, filter).Decode(result)
			if err != nil {
				return errors.Wrap(err)
			}
			return nil
		},
		InsertOne: func(ctx context.Context, collection string, doc any) (ResultInsertMongo, error) {
			res, err := d.client.Database(d.dbName).Collection(collection).InsertOne(ctx, doc)
			if err != nil {
				return ResultInsertMongo{}, errors.Wrap(err)
			}
			return ResultInsertMongo{InsertedID: res.InsertedID}, nil
		},
		InsertMany: func(ctx context.Context, collection string, docs []any) (ResultInsertManyMongo, error) {
			res, err := d.client.Database(d.dbName).Collection(collection).InsertMany(ctx, docs)
			if err != nil {
				return ResultInsertManyMongo{}, errors.Wrap(err)
			}
			return ResultInsertManyMongo{InsertedIDs: res.InsertedIDs}, nil
		},
		UpdateOne: func(ctx context.Context, collection string, filter any, update any) (ResultUpdateMongo, error) {
			res, err := d.client.Database(d.dbName).Collection(collection).UpdateOne(ctx, filter, update)
			if err != nil {
				return ResultUpdateMongo{}, errors.Wrap(err)
			}
			return ResultUpdateMongo{AffectedCount: res.ModifiedCount}, nil
		},
		UpdateMany: func(ctx context.Context, collection string, filter any, update any) (ResultUpdateMongo, error) {
			res, err := d.client.Database(d.dbName).Collection(collection).UpdateMany(ctx, filter, update)
			if err != nil {
				return ResultUpdateMongo{}, errors.Wrap(err)
			}
			return ResultUpdateMongo{AffectedCount: res.ModifiedCount}, nil
		},
		DeleteOne: func(ctx context.Context, collection string, filter any) (ResultDeleteMongo, error) {
			res, err := d.client.Database(d.dbName).Collection(collection).DeleteOne(ctx, filter)
			if err != nil {
				return ResultDeleteMongo{}, errors.Wrap(err)
			}
			return ResultDeleteMongo{AffectedCount: res.DeletedCount}, nil
		},
		DeleteMany: func(ctx context.Context, collection string, filter any) (ResultDeleteMongo, error) {
			res, err := d.client.Database(d.dbName).Collection(collection).DeleteMany(ctx, filter)
			if err != nil {
				return ResultDeleteMongo{}, errors.Wrap(err)
			}
			return ResultDeleteMongo{AffectedCount: res.DeletedCount}, nil
		},
	}
}
