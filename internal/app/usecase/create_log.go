package usecase

import (
	"context"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
)

type CreateLogInput struct {
	Level   string
	Message string
}

type CreateLogOutput struct {
	ID any
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
		return CreateLogOutput{}, err
	}

	return CreateLogOutput{ID: id}, nil
}
