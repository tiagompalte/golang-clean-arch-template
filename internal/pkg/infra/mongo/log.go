package mongo

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
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

func (r LogRepository) Find(ctx context.Context, filter any) ([]entity.Log, error) {
	result, err := r.conn.Find(ctx, r.collection, filter)
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
