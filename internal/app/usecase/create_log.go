package usecase

import (
	"context"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateLogInput struct {
	Level   string
	Message any
}

type CreateLogOutput struct {
	ID        string
	CreatedAt time.Time
	Level     string
	Message   any
}

type CreateLogUseCase interface {
	Execute(ctx context.Context, input CreateLogInput) (CreateLogOutput, error)
}

type CreateLogUseCaseImpl struct {
	logRepo protocols.LogRepository
}

func NewCreateLogUseCaseImpl(logRepo protocols.LogRepository) CreateLogUseCase {
	return CreateLogUseCaseImpl{
		logRepo: logRepo,
	}
}

func (uc CreateLogUseCaseImpl) Execute(ctx context.Context, input CreateLogInput) (CreateLogOutput, error) {
	var log entity.Log
	log.CreatedAt = time.Now()
	log.Level = input.Level
	log.Message = input.Message

	id, err := uc.logRepo.Insert(ctx, log)
	if err != nil {
		return CreateLogOutput{}, errors.Wrap(err)
	}

	objectID, ok := id.(bson.ObjectID)
	if !ok {
		return CreateLogOutput{}, errors.NewAppInternalServerError()
	}

	log, err = uc.logRepo.FindByID(ctx, objectID.Hex())
	if err != nil {
		return CreateLogOutput{}, errors.Wrap(err)
	}

	return CreateLogOutput{
		ID:        log.ID.Hex(),
		CreatedAt: log.CreatedAt,
		Level:     log.Level,
		Message:   log.Message,
	}, nil
}
