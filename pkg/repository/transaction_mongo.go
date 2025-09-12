package repository

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TransactionMongo struct {
	session   mongo.Session
	dbName    string
	committed bool
	rollback  bool
}

func newTransactionMongo(session mongo.Session, dbName string) TransactionMongoManager {
	transaction := new(TransactionMongo)
	transaction.session = session
	transaction.dbName = dbName
	return transaction
}

func (t *TransactionMongo) Commit() error {
	err := t.session.CommitTransaction(context.Background())
	if err != nil {
		return errors.Wrap(err)
	}

	t.committed = true

	return nil
}

func (t *TransactionMongo) Rollback() error {
	if t != nil && !t.committed && !t.rollback {
		err := t.session.AbortTransaction(context.Background())
		if err != nil {
			return errors.Wrap(err)
		}
		t.rollback = true
	}
	return nil
}

func (t *TransactionMongo) Command() ConnectorMongo {
	return ConnectorMongo{
		Find: func(ctx context.Context, collection string, filter any, opts ...options.Lister[options.FindOptions]) (RowsMongo, error) {
			cursor, err := t.session.Client().Database(t.dbName).Collection(collection).Find(ctx, filter, opts...)
			if err != nil {
				return nil, errors.Wrap(err)
			}
			return cursor, nil
		},
		InsertOne: func(ctx context.Context, collection string, doc any) (ResultInsertMongo, error) {
			res, err := t.session.Client().Database(t.dbName).Collection(collection).InsertOne(ctx, doc)
			if err != nil {
				return ResultInsertMongo{}, errors.Wrap(err)
			}
			return ResultInsertMongo{InsertedID: res.InsertedID}, nil
		},
		InsertMany: func(ctx context.Context, collection string, docs []any) (ResultInsertManyMongo, error) {
			res, err := t.session.Client().Database(t.dbName).Collection(collection).InsertMany(ctx, docs)
			if err != nil {
				return ResultInsertManyMongo{}, errors.Wrap(err)
			}
			return ResultInsertManyMongo{InsertedIDs: res.InsertedIDs}, nil
		},
		UpdateOne: func(ctx context.Context, collection string, filter any, update any) (ResultUpdateMongo, error) {
			res, err := t.session.Client().Database(t.dbName).Collection(collection).UpdateOne(ctx, filter, update)
			if err != nil {
				return ResultUpdateMongo{}, errors.Wrap(err)
			}
			return ResultUpdateMongo{AffectedCount: res.ModifiedCount}, nil
		},
		UpdateMany: func(ctx context.Context, collection string, filter any, update any) (ResultUpdateMongo, error) {
			res, err := t.session.Client().Database(t.dbName).Collection(collection).UpdateMany(ctx, filter, update)
			if err != nil {
				return ResultUpdateMongo{}, errors.Wrap(err)
			}
			return ResultUpdateMongo{AffectedCount: res.ModifiedCount}, nil
		},
		DeleteOne: func(ctx context.Context, collection string, filter any) (ResultDeleteMongo, error) {
			res, err := t.session.Client().Database(t.dbName).Collection(collection).DeleteOne(ctx, filter)
			if err != nil {
				return ResultDeleteMongo{}, errors.Wrap(err)
			}
			return ResultDeleteMongo{AffectedCount: res.DeletedCount}, nil
		},
		DeleteMany: func(ctx context.Context, collection string, filter any) (ResultDeleteMongo, error) {
			res, err := t.session.Client().Database(t.dbName).Collection(collection).DeleteMany(ctx, filter)
			if err != nil {
				return ResultDeleteMongo{}, errors.Wrap(err)
			}
			return ResultDeleteMongo{AffectedCount: res.DeletedCount}, nil
		},
	}
}
