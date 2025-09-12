package mongo

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LogRepository struct {
	conn       repository.ConnectorMongo
	collection string
}

func NewLogRepository(conn repository.ConnectorMongo) protocols.LogRepository {
	return LogRepository{
		conn:       conn,
		collection: "logs",
	}
}

func (r LogRepository) Insert(ctx context.Context, log entity.Log) (any, error) {
	res, err := r.conn.InsertOne(ctx, r.collection, log)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return res.InsertedID, nil
}

func (r LogRepository) FindAll(ctx context.Context, limit int64) ([]entity.Log, error) {
	result, err := r.conn.Find(ctx, r.collection, bson.D{}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var logs []entity.Log
	for result.Next(ctx) {
		var log entity.Log
		if err := result.Decode(&log); err != nil {
			return nil, errors.Wrap(err)
		}
		logs = append(logs, log)
	}

	return logs, nil
}

func (r LogRepository) FindByID(ctx context.Context, id string) (entity.Log, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return entity.Log{}, errors.Wrap(err)
	}

	var log entity.Log
	err = r.conn.FindOne(ctx, r.collection, bson.M{"_id": objectID}, &log)
	if err != nil {
		return entity.Log{}, errors.Wrap(err)
	}
	return log, nil
}
